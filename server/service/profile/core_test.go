package profile

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

// config

var queries *repo.Queries
var conn *sql.DB

func init() {
	conn = db.GetConnection(":memory:")
	db.ConfigDatabase(conn, "../../db/schema.sql")
	queries = repo.New(conn)
}

func configCtx() context.Context {
	ctx := context.Background()
	ctx = db.AttachQueries(ctx, queries)
	return ctx
}

func seedProfile(t *testing.T, id string) {
	t.Helper()

	email := id + "@test.com"
	username := id + "_username"

	_, err := conn.Exec(`
		INSERT INTO profile
			(id, email, username, displayname, color, password)
		VALUES
			(?, ?, ?, 'test-displayname', '#000000', 'test-password')`,
		id, email, username)
	if err != nil {
		t.Fatalf("failed to seed profile: %v", err)
	}

	t.Cleanup(func() {
		if _, err := conn.Exec("DELETE FROM profile WHERE id = ?", id); err != nil {
			t.Fatalf("failed to cleanup test profile: %v", err)
		}
	})
}

func seedPost(t *testing.T, userID, postID string, createdAt time.Time, visibility types.Visibility) {
	t.Helper()

	_, err := conn.Exec(`
		INSERT INTO post
			(id, user_id, created_at, updated_at, visibility, word, definition, part_of_speech, pronunciation, example, synonyms, antonyms)
		VALUES
			(?, ?, ?, ?, ?, 'test-word', 'test-definition', 'test-part-of-speech', 'test-pronunciation', 'test-example', '[]', '[]')`,
		postID, userID, createdAt, createdAt, string(visibility),
	)
	if err != nil {
		t.Fatalf("failed to seed post: %v", err)
	}
}

func asInt64(t *testing.T, v any) int64 {
	t.Helper()

	switch x := v.(type) {
	case int64:
		return x
	case int:
		return int64(x)
	case int32:
		return int64(x)
	case float64:
		return int64(x)
	case []byte:
		// SQLite sometimes returns numeric columns as []byte.
		i, err := strconv.ParseInt(string(x), 10, 64)
		if err != nil {
			t.Fatalf("failed to parse []byte int64: %v", err)
		}
		return i
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			t.Fatalf("failed to parse string int64: %v", err)
		}
		return i
	case nil:
		return 0
	default:
		t.Fatalf("unexpected type for numeric interface: %T", v)
		return 0
	}
}

// testing

func TestGetProfileCountsAndIsFollowing(t *testing.T) {
	ctx := configCtx()

	followerID := "test-follower-id"
	followeeID := "test-followee-id"

	seedProfile(t, followerID)
	seedProfile(t, followeeID)

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 2; i++ {
		seedPost(t, followeeID, "followee-post-"+strconv.Itoa(i), base.Add(time.Second*time.Duration(i)), types.Public)
	}
	seedPost(t, followerID, "follower-post-0", base.Add(time.Second*3), types.Public)

	t.Run("not following", func(t *testing.T) {
		profile, apiErr := getProfile(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if profile.PostCount != 2 {
			t.Fatalf("expected followee post count 2, got %v", profile.PostCount)
		}
		if profile.FollowingCount != 0 {
			t.Fatalf("expected followee following count 0, got %v", profile.FollowingCount)
		}
		if profile.FollowerCount != 0 {
			t.Fatalf("expected followee follower count 0, got %v", profile.FollowerCount)
		}
		if isFollowing := asInt64(t, profile.IsFollowing); isFollowing != 0 {
			t.Fatalf("expected isFollowing false, got %v", isFollowing)
		}
	})

	t.Run("after follow", func(t *testing.T) {
		apiErr := followUser(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		profile, apiErr := getProfile(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if profile.PostCount != 2 {
			t.Fatalf("expected followee post count 2, got %v", profile.PostCount)
		}
		if profile.FollowingCount != 1 {
			t.Fatalf("expected followee following count 1, got %v", profile.FollowingCount)
		}
		if profile.FollowerCount != 0 {
			t.Fatalf("expected followee follower count 0, got %v", profile.FollowerCount)
		}
		if isFollowing := asInt64(t, profile.IsFollowing); isFollowing != 1 {
			t.Fatalf("expected isFollowing true, got %v", isFollowing)
		}

		// From the followee's perspective, they are not following the follower.
		profileForFollower, apiErr := getProfile(ctx, followeeID, followerID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if profileForFollower.FollowerCount != 1 {
			t.Fatalf("expected follower follower count 1, got %v", profileForFollower.FollowerCount)
		}
		if profileForFollower.FollowingCount != 0 {
			t.Fatalf("expected follower following count 0, got %v", profileForFollower.FollowingCount)
		}
		if isFollowing := asInt64(t, profileForFollower.IsFollowing); isFollowing != 0 {
			t.Fatalf("expected isFollowing false, got %v", isFollowing)
		}
	})
}

func TestGetProfilePostsVisibilityAndCounts(t *testing.T) {
	ctx := configCtx()

	viewerID := "test-viewer-id"
	profileID := "test-profile-id"

	seedProfile(t, viewerID)
	seedProfile(t, profileID)

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	// Newer private posts that the viewer should not see.
	for i := 0; i < 5; i++ {
		seedPost(t, profileID, "private-post-"+strconv.Itoa(i), base.Add(time.Second*time.Duration(10+i)), types.Private)
	}
	// Older public posts that the viewer should see.
	for i := 0; i < 5; i++ {
		seedPost(t, profileID, "public-post-"+strconv.Itoa(i), base.Add(time.Second*time.Duration(i)), types.Public)
	}

	t.Run("viewer cannot see private", func(t *testing.T) {
		rows, apiErr := getProfilePosts(ctx, viewerID, profileID, 0)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if len(rows) != 5 {
			t.Fatalf("expected 5 rows, got %v", len(rows))
		}
		if rows[0].ID != "public-post-4" {
			t.Fatalf("expected newest public post id %v, got %v", "public-post-4", rows[0].ID)
		}
		for _, r := range rows {
			if r.Visibility != string(types.Public) {
				t.Fatalf("expected row visibility public, got %v", r.Visibility)
			}
		}

		total, apiErr := getProfilePostsCount(ctx, viewerID, profileID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if total != 5 {
			t.Fatalf("expected posts count 5, got %v", total)
		}
	})

	t.Run("owner can see private", func(t *testing.T) {
		rows, apiErr := getProfilePosts(ctx, profileID, profileID, 0)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if len(rows) != 10 {
			t.Fatalf("expected 10 rows, got %v", len(rows))
		}
		if rows[0].ID != "private-post-4" {
			t.Fatalf("expected newest private post id %v, got %v", "private-post-4", rows[0].ID)
		}

		total, apiErr := getProfilePostsCount(ctx, profileID, profileID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if total != 10 {
			t.Fatalf("expected posts count 10, got %v", total)
		}
	})
}

func TestGetProfilePostsPagination(t *testing.T) {
	ctx := configCtx()

	viewerID := "test-viewer-pagination-id"
	profileID := "test-profile-pagination-id"

	seedProfile(t, viewerID)
	seedProfile(t, profileID)

	totalPosts := config.Env.PostsLimit + 5

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < totalPosts; i++ {
		seedPost(
			t,
			profileID,
			"public-pagination-post-"+strconv.Itoa(i),
			base.Add(time.Second*time.Duration(i)),
			types.Public,
		)
	}

	t.Run("page 0 returns latest", func(t *testing.T) {
		rows, apiErr := getProfilePosts(ctx, viewerID, profileID, 0)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if len(rows) != config.Env.PostsLimit {
			t.Fatalf("expected %v rows, got %v", config.Env.PostsLimit, len(rows))
		}

		expectedFirstIdx := totalPosts - 1
		if rows[0].ID != "public-pagination-post-"+strconv.Itoa(expectedFirstIdx) {
			t.Fatalf("expected newest id %v, got %v", expectedFirstIdx, rows[0].ID)
		}
	})

	t.Run("page 1 returns remaining", func(t *testing.T) {
		if totalPosts <= config.Env.PostsLimit {
			t.Skip("not enough posts to paginate")
		}

		rows, apiErr := getProfilePosts(ctx, viewerID, profileID, 1)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		expectedLen := totalPosts - config.Env.PostsLimit
		if len(rows) != expectedLen {
			t.Fatalf("expected %v rows, got %v", expectedLen, len(rows))
		}

		expectedFirstIdx := totalPosts - 1 - config.Env.PostsLimit
		if rows[0].ID != "public-pagination-post-"+strconv.Itoa(expectedFirstIdx) {
			t.Fatalf("expected newest page-1 id %v, got %v", expectedFirstIdx, rows[0].ID)
		}
	})
}

func TestFollowUserAndUnfollowUser(t *testing.T) {
	ctx := configCtx()

	followerID := "test-follow-unfollow-follower-id"
	followeeID := "test-follow-unfollow-followee-id"

	seedProfile(t, followerID)
	seedProfile(t, followeeID)

	t.Run("follow updates isFollowing and counts", func(t *testing.T) {
		before, apiErr := getProfile(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if isFollowing := asInt64(t, before.IsFollowing); isFollowing != 0 {
			t.Fatalf("expected not following, got %v", isFollowing)
		}
		if before.FollowingCount != 0 {
			t.Fatalf("expected following count 0, got %v", before.FollowingCount)
		}

		apiErr = followUser(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		after, apiErr := getProfile(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if isFollowing := asInt64(t, after.IsFollowing); isFollowing != 1 {
			t.Fatalf("expected isFollowing true, got %v", isFollowing)
		}
		if after.FollowingCount != 1 {
			t.Fatalf("expected following count 1, got %v", after.FollowingCount)
		}
	})

	t.Run("unfollow resets isFollowing and counts", func(t *testing.T) {
		apiErr := unfollowUser(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		after, apiErr := getProfile(ctx, followerID, followeeID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if isFollowing := asInt64(t, after.IsFollowing); isFollowing != 0 {
			t.Fatalf("expected isFollowing false, got %v", isFollowing)
		}
		if after.FollowingCount != 0 {
			t.Fatalf("expected following count 0, got %v", after.FollowingCount)
		}
	})
}

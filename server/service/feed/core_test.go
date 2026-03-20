package feed

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

func seedUser(t *testing.T, id string) {
	t.Helper()

	_, err := conn.Exec(`
	INSERT INTO profile
		(id, email, username, displayname, color, password)
	VALUES
		($1, $1, $1, 'test-displayname', '#000000', 'test-password')`,
		id)
	if err != nil {
		t.Fatalf("failed to seed profile: %v", err)
	}

	t.Cleanup(func() {
		if _, err := conn.Exec("DELETE FROM profile WHERE id = ?", id); err != nil {
			t.Fatalf("failed to cleanup test user: %v", err)
		}
	})
}

func seedPost(t *testing.T, userID, postID string, idx int) {
	t.Helper()

	_, err := conn.Exec(`
		INSERT INTO post
			(id, user_id, created_at, visibility, word, definition, part_of_speech, pronunciation, example, synonyms, antonyms)
		VALUES
			(?, ?, ?, 'public', 'test-word', 'test-definition', 'test-part-of-speech', 'test-pronunciation', 'test-example', '[]', '[]')
	`, postID, userID, time.Now().Add(time.Second*time.Duration(idx)))
	if err != nil {
		t.Fatalf("failed to seed post: %v", err)
	}
}

// testing

func TestGetRecentFeed(t *testing.T) {
	ctx := configCtx()

	userID := "test-user-id"
	seedUser(t, userID)
	for i := range 100 {
		seedPost(t, userID, "test-post-id-"+strconv.Itoa(i), i)
	}

	t.Run("success", func(t *testing.T) {
		rows, apiErr := getRecentPosts(ctx, userID, 0)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if len(rows) != config.Env.PostsLimit {
			t.Errorf("expected %v rows, got %v", config.Env.PostsLimit, len(rows))
		}

		if rows[0].ID != "test-post-id-99" {
			t.Errorf("expected %v, got %v", "test-post-id-99", rows[0].ID)
		}
	})

	t.Run("page 1 uses offset", func(t *testing.T) {
		if 100 <= config.Env.PostsLimit {
			t.Skip("not enough seeded posts for pagination test")
		}

		rows, apiErr := getRecentPosts(ctx, userID, 1)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if len(rows) != config.Env.PostsLimit {
			t.Errorf("expected %v rows, got %v", config.Env.PostsLimit, len(rows))
		}

		expectedFirstID := "test-post-id-" + strconv.Itoa(99-config.Env.PostsLimit)
		expectedLastID := "test-post-id-" + strconv.Itoa(99-(2*config.Env.PostsLimit-1))

		if rows[0].ID != expectedFirstID {
			t.Errorf("expected %v, got %v", expectedFirstID, rows[0].ID)
		}
		if rows[len(rows)-1].ID != expectedLastID {
			t.Errorf("expected %v, got %v", expectedLastID, rows[len(rows)-1].ID)
		}
	})
}

func TestGetRecentPostsCount(t *testing.T) {
	ctx := configCtx()

	userID := "test-user-id"
	postsCount := 100
	seedUser(t, userID)
	for i := range postsCount {
		seedPost(t, userID, "test-post-id-"+strconv.Itoa(i), i)
	}

	t.Run("success", func(t *testing.T) {
		total, apiErr := getRecentPostsCount(ctx)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if total != postsCount {
			t.Errorf("expected %v, got %v", postsCount, total)
		}
	})
}

func TestGetFollowingFeed(t *testing.T) {
	ctx := configCtx()

	userID := "test-user-id"
	seedUser(t, userID)
	for i := range 100 {
		seedPost(t, userID, "test-post-id-"+strconv.Itoa(i), i)
	}

	followerID := "test-following-user-id"
	seedUser(t, followerID)

	_, err := conn.Exec(`
		INSERT INTO user_follows
			(follower_id, followee_id)
		VALUES
			(?, ?)`,
		followerID, userID)
	if err != nil {
		t.Fatalf("failed to seed follower: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		rows, apiErr := getFollowingPosts(ctx, followerID, 0)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if len(rows) != config.Env.PostsLimit {
			t.Errorf("expected %v rows, got %v", config.Env.PostsLimit, len(rows))
		}

		if rows[0].ID != "test-post-id-99" {
			t.Errorf("expected %v, got %v", "test-post-id-99", rows[0].ID)
		}
	})

	t.Run("page 1 uses offset", func(t *testing.T) {
		if 100 <= config.Env.PostsLimit {
			t.Skip("not enough seeded posts for pagination test")
		}

		rows, apiErr := getFollowingPosts(ctx, followerID, 1)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if len(rows) != config.Env.PostsLimit {
			t.Errorf("expected %v rows, got %v", config.Env.PostsLimit, len(rows))
		}

		expectedFirstID := "test-post-id-" + strconv.Itoa(99-config.Env.PostsLimit)
		expectedLastID := "test-post-id-" + strconv.Itoa(99-(2*config.Env.PostsLimit-1))

		if rows[0].ID != expectedFirstID {
			t.Errorf("expected %v, got %v", expectedFirstID, rows[0].ID)
		}
		if rows[len(rows)-1].ID != expectedLastID {
			t.Errorf("expected %v, got %v", expectedLastID, rows[len(rows)-1].ID)
		}
	})
}

func TestGetFollowingPostsCount(t *testing.T) {
	ctx := configCtx()

	userID := "test-user-id"
	postsCount := 100
	seedUser(t, userID)
	for i := range postsCount {
		seedPost(t, userID, "test-post-id-"+strconv.Itoa(i), i)
	}

	followerID := "test-following-user-id"
	seedUser(t, followerID)

	_, err := conn.Exec(`
		INSERT INTO user_follows
			(follower_id, followee_id)
		VALUES
			(?, ?)`,
		followerID, userID)
	if err != nil {
		t.Fatalf("failed to seed follower: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		total, apiErr := getFollowingPostsCount(ctx, followerID)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if total != postsCount {
			t.Errorf("expected %v, got %v", postsCount, total)
		}
	})
}

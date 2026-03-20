package post

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"testing"
	"time"

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

func countRows(t *testing.T, query string, args ...any) int64 {
	t.Helper()

	var count int64
	if err := conn.QueryRow(query, args...).Scan(&count); err != nil {
		t.Fatalf("failed to count rows: %v", err)
	}
	return count
}

func TestCreatePost(t *testing.T) {
	ctx := configCtx()

	userID := "test-create-post-user-id"
	seedProfile(t, userID)

	reqJSON := `{
		"visibility": "public",
		"content": {
			"word": "test-word-1",
			"definition": "test-definition-1",
			"partOfSpeech": "noun",
			"pronunciation": "test-pronunciation-1",
			"synonyms": ["syn1", "syn2"],
			"antonyms": ["ant1"],
			"example": "test-example-1"
		}
	}`

	var req types.CreatePostRequest
	if err := json.Unmarshal([]byte(reqJSON), &req); err != nil {
		t.Fatalf("failed to unmarshal test request: %v", err)
	}

	apiErr := createPost(ctx, userID, req)
	if apiErr != types.Errors.None {
		t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
	}

	var insertedID string
	var insertedVisibility string
	var insertedWord string
	var insertedSynonyms string
	var insertedAntonyms string

	if err := conn.QueryRow(`
		SELECT id, visibility, word, synonyms, antonyms
		FROM post
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT 1`,
		userID,
	).Scan(&insertedID, &insertedVisibility, &insertedWord, &insertedSynonyms, &insertedAntonyms); err != nil {
		t.Fatalf("failed to query inserted post: %v", err)
	}

	if insertedVisibility != string(types.Public) {
		t.Fatalf("expected visibility %v, got %v", types.Public, insertedVisibility)
	}
	if insertedWord != "test-word-1" {
		t.Fatalf("expected word %v, got %v", "test-word-1", insertedWord)
	}
	if insertedSynonyms != `["syn1","syn2"]` {
		t.Fatalf("expected synonyms %v, got %v", `["syn1","syn2"]`, insertedSynonyms)
	}
	if insertedAntonyms != `["ant1"]` {
		t.Fatalf("expected antonyms %v, got %v", `["ant1"]`, insertedAntonyms)
	}
	if insertedID == "" {
		t.Fatalf("expected inserted id to be non-empty")
	}
}

func TestDeletePost(t *testing.T) {
	ctx := configCtx()

	userID := "test-delete-post-user-id"
	postID := "test-delete-post-id"

	seedProfile(t, userID)

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	seedPost(t, userID, postID, base, types.Public)

	apiErr := deletePost(ctx, postID)
	if apiErr != types.Errors.None {
		t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
	}

	if count := countRows(t, `SELECT COUNT(*) FROM post WHERE id = ?`, postID); count != 0 {
		t.Fatalf("expected deleted post count 0, got %v", count)
	}

	// Deleting a non-existent post should be a no-op.
	apiErr = deletePost(ctx, "non-existent-post-id-" + strconv.Itoa(int(time.Now().UnixNano())))
	if apiErr != types.Errors.None {
		t.Fatalf("expected nil apiErr when deleting missing post, got %v", apiErr.WrappedError)
	}
}

func TestLikeUnlikePost(t *testing.T) {
	ctx := configCtx()

	ownerID := "test-like-owner-id"
	likerID := "test-like-user-id"
	postID := "test-like-post-id"

	seedProfile(t, ownerID)
	seedProfile(t, likerID)

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	seedPost(t, ownerID, postID, base, types.Public)

	t.Run("like creates row", func(t *testing.T) {
		apiErr := likePost(ctx, likerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if count := countRows(t, `SELECT COUNT(*) FROM post_like WHERE user_id = ? AND post_id = ?`, likerID, postID); count != 1 {
			t.Fatalf("expected like row count 1, got %v", count)
		}
	})

	t.Run("liking twice fails with internal server error", func(t *testing.T) {
		apiErr := likePost(ctx, likerID, postID)
		if apiErr == types.Errors.None {
			t.Fatalf("expected error when liking twice")
		}
		if apiErr.Code != types.Errors.InternalServerError(nil).Code {
			t.Fatalf("expected InternalServerError code, got %v", apiErr.Code)
		}
		if count := countRows(t, `SELECT COUNT(*) FROM post_like WHERE user_id = ? AND post_id = ?`, likerID, postID); count != 1 {
			t.Fatalf("expected like row count to remain 1, got %v", count)
		}
	})

	t.Run("unlike deletes row (and is idempotent)", func(t *testing.T) {
		apiErr := unlikePost(ctx, likerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if count := countRows(t, `SELECT COUNT(*) FROM post_like WHERE user_id = ? AND post_id = ?`, likerID, postID); count != 0 {
			t.Fatalf("expected like row count 0, got %v", count)
		}

		apiErr = unlikePost(ctx, likerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr on unlike missing row, got %v", apiErr.WrappedError)
		}
	})
}

func TestBookmarkUnbookmarkPost(t *testing.T) {
	ctx := configCtx()

	ownerID := "test-bookmark-owner-id"
	bookmarkerID := "test-bookmark-user-id"
	postID := "test-bookmark-post-id"

	seedProfile(t, ownerID)
	seedProfile(t, bookmarkerID)

	base := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	seedPost(t, ownerID, postID, base, types.Public)

	t.Run("bookmark creates row", func(t *testing.T) {
		apiErr := bookmarkPost(ctx, bookmarkerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}

		if count := countRows(t, `SELECT COUNT(*) FROM post_bookmark WHERE user_id = ? AND post_id = ?`, bookmarkerID, postID); count != 1 {
			t.Fatalf("expected bookmark row count 1, got %v", count)
		}
	})

	t.Run("bookmarking twice fails with internal server error", func(t *testing.T) {
		apiErr := bookmarkPost(ctx, bookmarkerID, postID)
		if apiErr == types.Errors.None {
			t.Fatalf("expected error when bookmarking twice")
		}
		if apiErr.Code != types.Errors.InternalServerError(nil).Code {
			t.Fatalf("expected InternalServerError code, got %v", apiErr.Code)
		}
		if count := countRows(t, `SELECT COUNT(*) FROM post_bookmark WHERE user_id = ? AND post_id = ?`, bookmarkerID, postID); count != 1 {
			t.Fatalf("expected bookmark row count to remain 1, got %v", count)
		}
	})

	t.Run("unbookmark deletes row (and is idempotent)", func(t *testing.T) {
		apiErr := unbookmarkPost(ctx, bookmarkerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr, got %v", apiErr.WrappedError)
		}
		if count := countRows(t, `SELECT COUNT(*) FROM post_bookmark WHERE user_id = ? AND post_id = ?`, bookmarkerID, postID); count != 0 {
			t.Fatalf("expected bookmark row count 0, got %v", count)
		}

		apiErr = unbookmarkPost(ctx, bookmarkerID, postID)
		if apiErr != types.Errors.None {
			t.Fatalf("expected nil apiErr on unbookmark missing row, got %v", apiErr.WrappedError)
		}
	})
}


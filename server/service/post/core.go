package post

import (
	"context"
	"encoding/json"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

func createPost(ctx context.Context, userID string, body types.CreatePostRequest) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	synonyms, err := json.Marshal(body.Content.Synonyms)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}
	antonyms, err := json.Marshal(body.Content.Antonyms)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	_, err = q.CreatePost(ctx, repo.CreatePostParams{
		Word:          body.Content.Word,
		Definition:    body.Content.Definition,
		PartOfSpeech:  body.Content.PartOfSpeech,
		Pronunciation: body.Content.Pronunciation,
		Synonyms:      string(synonyms),
		Antonyms:      string(antonyms),
		Example:       body.Content.Example,
		Visibility:    body.Visibility,
		UserID:        userID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func deletePost(ctx context.Context, postID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.DeletePost(ctx, postID)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func likePost(ctx context.Context, userID, postID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.LikePost(ctx, repo.LikePostParams{
		UserID: userID,
		PostID: postID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func unlikePost(ctx context.Context, userID, postID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.UnlikePost(ctx, repo.UnlikePostParams{
		UserID: userID,
		PostID: postID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func bookmarkPost(ctx context.Context, userID, postID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.BookmarkPost(ctx, repo.BookmarkPostParams{
		UserID: userID,
		PostID: postID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func unbookmarkPost(ctx context.Context, userID, postID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.UnbookmarkPost(ctx, repo.UnbookmarkPostParams{
		UserID: userID,
		PostID: postID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

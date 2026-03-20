package feed

import (
	"context"

	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

var limit = config.Env.PostsLimit

func getRecentPosts(ctx context.Context, userID string, page int) ([]repo.GetRecentPostsRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	rows, err := q.GetRecentPosts(ctx, repo.GetRecentPostsParams{
		UserID: userID,
		Limit:  int64(limit),
		Offset: int64(limit * page),
	})
	if err != nil {
		return nil, types.Errors.InternalServerError(err)
	}

	return rows, types.Errors.None
}

func getRecentPostsCount(ctx context.Context) (int, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	count, err := q.GetTotalPostCount(ctx)
	if err != nil {
		return 0, types.Errors.InternalServerError(err)
	}

	return int(count), types.Errors.None
}

func getFollowingPosts(ctx context.Context, userID string, page int) ([]repo.GetFollowingPostsRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	rows, err := q.GetFollowingPosts(ctx, repo.GetFollowingPostsParams{
		UserID: userID,
		Limit:  int64(limit),
		Offset: int64(limit * page),
	})
	if err != nil {
		return nil, types.Errors.InternalServerError(err)
	}

	return rows, types.Errors.None
}

func getFollowingPostsCount(ctx context.Context, userID string) (int, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	count, err := q.GetTotalFollowingPostCount(ctx, userID)
	if err != nil {
		return 0, types.Errors.InternalServerError(err)
	}

	return int(count), types.Errors.None
}

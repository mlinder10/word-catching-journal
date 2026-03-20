package profile

import (
	"context"

	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

var postLimit = config.Env.PostsLimit

func getProfile(ctx context.Context, userID, profileID string) (repo.GetProfileRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.GetProfile(ctx, repo.GetProfileParams{
		UserID:    userID,
		ProfileID: profileID,
	})
	if err != nil {
		return repo.GetProfileRow{}, types.Errors.InternalServerError(err)
	}

	return row, types.Errors.None
}

func getProfilePosts(ctx context.Context, userID, profileID string, page int) ([]repo.GetProfilePostsRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	rows, err := q.GetProfilePosts(ctx, repo.GetProfilePostsParams{
		UserID:        userID,
		ProfileID:     profileID,
		CanSeePrivate: userID == profileID,
		Limit:         int64(postLimit),
		Offset:        int64(postLimit * page),
	})
	if err != nil {
		return nil, types.Errors.InternalServerError(err)
	}

	return rows, types.Errors.None
}

func getProfilePostsCount(ctx context.Context, userID, profileID string) (int, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	count, err := q.GetTotalProfilePostCount(ctx, repo.GetTotalProfilePostCountParams{
		ProfileID:     profileID,
		CanSeePrivate: userID == profileID,
	})
	if err != nil {
		return 0, types.Errors.InternalServerError(err)
	}

	return int(count), types.Errors.None
}

func followUser(ctx context.Context, userID, profileID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.FollowUser(ctx, repo.FollowUserParams{
		FollowerID: userID,
		FolloweeID: profileID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func unfollowUser(ctx context.Context, userID, profileID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.UnfollowUser(ctx, repo.UnfollowUserParams{
		FollowerID: userID,
		FolloweeID: profileID,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

var userLimit = config.Env.UsersLimit

func getFollowing(ctx context.Context, userID string, page int) ([]repo.GetFollowingRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	rows, err := q.GetFollowing(ctx, repo.GetFollowingParams{
		UserID: userID,
		Limit:  int64(userLimit),
		Offset: int64(page * userLimit),
	})
	if err != nil {
		return nil, types.Errors.InternalServerError(err)
	}

	return rows, types.Errors.None
}

func getFollowers(ctx context.Context, userID string, page int) ([]repo.GetFollowersRow, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	rows, err := q.GetFollowers(ctx, repo.GetFollowersParams{
		UserID: userID,
		Limit:  int64(userLimit),
		Offset: int64(page * userLimit),
	})
	if err != nil {
		return nil, types.Errors.InternalServerError(err)
	}

	return rows, types.Errors.None
}

func getFollowingCount(ctx context.Context, userID string) (int, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	count, err := q.GetFollowingCount(ctx, userID)
	if err != nil {
		return 0, types.Errors.InternalServerError(err)
	}

	return int(count), types.Errors.None
}

func getFollowersCount(ctx context.Context, userID string) (int, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	count, err := q.GetFollowersCount(ctx, userID)
	if err != nil {
		return 0, types.Errors.InternalServerError(err)
	}

	return int(count), types.Errors.None
}

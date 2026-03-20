package post

import (
	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/types"
)

var limit = config.Env.PostsLimit

func BuildResponse[T types.PostMapable](
	posts []T, page, totalPosts int,
) (types.PostResponse, types.APIError) {
	parsedPosts := make([]types.Post, len(posts))

	if len(posts) == 0 {
		return types.PostResponse{
			Posts:       []types.Post{},
			CurrentPage: page + 1,
			TotalPages:  page + 1,
		}, types.Errors.None
	}

	for i, post := range posts {
		parsedPosts[i] = post.ToPost()
	}

	return types.PostResponse{
		Posts:       parsedPosts,
		CurrentPage: page + 1,
		TotalPages:  totalPosts/limit + 1,
	}, types.Errors.None
}

package feed

import (
	"net/http"

	"github.com/mlinder10/wcj/service/auth"
	"github.com/mlinder10/wcj/service/post"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func GET_FeedRecent(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	page := utils.ParsePageQuery(r)

	posts, apiErr := getRecentPosts(r.Context(), userID, page)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}
	count, apiErr := getRecentPostsCount(r.Context())
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	response, apiErr := post.BuildResponse(posts, page, count)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, response)
}

func GET_FeedFollowing(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	page := utils.ParsePageQuery(r)

	posts, apiErr := getFollowingPosts(r.Context(), userID, page)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	count, apiErr := getFollowingPostsCount(r.Context(), userID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	response, apiErr := post.BuildResponse(posts, page, count)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, response)
}

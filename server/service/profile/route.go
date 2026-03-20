package profile

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlinder10/wcj/service/auth"
	"github.com/mlinder10/wcj/service/post"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func GET_ProfileID(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	profileID := mux.Vars(r)["profileID"]

	profile, apiErr := getProfile(r.Context(), userID, profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, profile)
}

func GET_ProfileIDPosts(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	profileID := mux.Vars(r)["profileID"]
	page := utils.ParsePageQuery(r)

	posts, apiErr := getProfilePosts(r.Context(), userID, profileID, page)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	count, apiErr := getProfilePostsCount(r.Context(), userID, profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	response, err := post.BuildResponse(posts, page, count)
	if err != types.Errors.None {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, 200, response)
}

func POST_ProfileIDFollow(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	profileID := mux.Vars(r)["profileID"]

	apiErr := followUser(r.Context(), userID, profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func DELETE_ProfileIDFollow(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	profileID := mux.Vars(r)["profileID"]

	apiErr := unfollowUser(r.Context(), userID, profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func GET_ProfileIDFollowing(w http.ResponseWriter, r *http.Request) {
	profileID := mux.Vars(r)["profileID"]
	page := utils.ParsePageQuery(r)

	rows, apiErr := getFollowing(r.Context(), profileID, page)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	count, apiErr := getFollowingCount(r.Context(), profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	response, apiErr := BuildResponse(rows, page, count)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, response)
}

func GET_ProfileIDFollowers(w http.ResponseWriter, r *http.Request) {
	profileID := mux.Vars(r)["profileID"]
	page := utils.ParsePageQuery(r)

	rows, apiErr := getFollowers(r.Context(), profileID, page)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	count, apiErr := getFollowersCount(r.Context(), profileID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	response, apiErr := BuildResponse(rows, page, count)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, response)
}

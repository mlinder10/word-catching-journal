package post

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlinder10/wcj/service/auth"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func POST_Post(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	var body types.CreatePostRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	apiErr := createPost(r.Context(), userID, body)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func DELETE_PostID(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["postID"]

	apiErr := deletePost(r.Context(), postID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func POST_PostIDLike(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	postID := mux.Vars(r)["postID"]

	apiErr := likePost(r.Context(), userID, postID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func DELETE_PostIDLike(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	postID := mux.Vars(r)["postID"]

	apiErr := unlikePost(r.Context(), userID, postID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func POST_PostIDBookmark(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	postID := mux.Vars(r)["postID"]

	apiErr := bookmarkPost(r.Context(), userID, postID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func DELETE_PostIDBookmark(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	postID := mux.Vars(r)["postID"]

	apiErr := unbookmarkPost(r.Context(), userID, postID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

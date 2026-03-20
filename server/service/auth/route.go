package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/email"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func POST_Login(w http.ResponseWriter, r *http.Request) {
	var body types.LoginRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	user, apiErr := loginUser(r.Context(), body.Email, body.Password)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	token, err := createJWT([]byte(config.Env.JWTSecret), user.ID)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	setJWTCookie(w, token)
	utils.WriteJSON(w, 200, user)
}

func POST_Register(w http.ResponseWriter, r *http.Request) {
	var body types.RegisterRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	apiErr := deleteUnverifiedProfile(r.Context(), body.Email, body.Username)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	user, apiErr := registerUser(r.Context(), body.Email, body.Username, body.Displayname, body.Password)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	code := utils.ConfirmationCode()
	if apiErr = createVerificationRecord(r.Context(), user.ID, code); apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	err = email.SendConfirmation(user.Email, code)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	utils.WriteJSON(w, 200, user)
}

func GET_Refresh(w http.ResponseWriter, r *http.Request) {
	userID := GetUserIDFromContext(r.Context())

	user, apiErr := verifyToken(r.Context(), userID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	token, err := createJWT([]byte(config.Env.JWTSecret), user.ID)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	setJWTCookie(w, token)
	utils.WriteJSON(w, 200, user)
}

func POST_Logout(w http.ResponseWriter, r *http.Request) {
	deleteJWTCookie(w)
	utils.WriteJSON(w, 200, nil)
}

func POST_Verify(w http.ResponseWriter, r *http.Request) {
	var body types.VerifyEmailRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	user, apiErr := verifyEmail(r.Context(), body.Code)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	token, err := createJWT([]byte(config.Env.JWTSecret), user.ID)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	setJWTCookie(w, token)
	utils.WriteJSON(w, 200, user)
}

func POST_VerifyResendID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	apiErr := deleteVerificationRecord(r.Context(), userID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	code := utils.ConfirmationCode()
	if apiErr := createVerificationRecord(r.Context(), userID, code); apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	user, apiErr := getUserById(r.Context(), userID)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	err := email.SendConfirmation(user.Email, code)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func PATCH_ResetPassword(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	var body types.ChangePasswordRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	apiErr := changePassword(r.Context(), userID, body.OldPassword, body.NewPassword)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func POST_ResetPassword(w http.ResponseWriter, r *http.Request) {
	var body types.ResetPasswordRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	code, apiErr := createPasswordResetRecord(r.Context(), body.Email)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	err = email.SendResetPassword(body.Email, code)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	utils.WriteJSON(w, 200, nil)
}

func POST_ResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	var body types.UpdatePasswordRequest
	err := utils.ParseJSON(r, &body)
	if err != nil {
		utils.WriteError(w, types.Errors.FailedToParseRequest(err))
		return
	}

	apiErr := resetPassword(r.Context(), code, body.Password)
	if apiErr != types.Errors.None {
		utils.WriteError(w, apiErr)
		return
	}

	utils.WriteJSON(w, 200, nil)
}

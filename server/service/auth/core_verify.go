package auth

import (
	"context"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

func createVerificationRecord(ctx context.Context, userID, code string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.CreateVerificationRecord(ctx, repo.CreateVerificationRecordParams{
		UserID: userID,
		Code:   code,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func verifyEmail(ctx context.Context, code string) (types.UserResponse, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	verificationRow, err := q.GetVerificationByCode(ctx, code)
	if err != nil {
		return types.UserResponse{}, types.Errors.InternalServerError(err)
	}

	row, err := q.VerifyUser(ctx, verificationRow.UserID)
	if err != nil {
		return types.UserResponse{}, types.Errors.InternalServerError(err)
	}

	user := types.UserResponse{
		ID:          row.ID,
		Email:       row.Email,
		Username:    row.Username,
		Displayname: row.Displayname,
		Color:       row.Color,
		ImageURL:    row.ImageUrl,
		CreatedAt:   row.CreatedAt,
	}

	return user, types.Errors.None
}

func deleteVerificationRecord(ctx context.Context, userID string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.DeleteVerificationRecord(ctx, userID)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

func deleteUnverifiedProfile(ctx context.Context, email, username string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	err := q.DeleteUnverifiedAccount(ctx, repo.DeleteUnverifiedAccountParams{
		Email:    email,
		Username: username,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

package auth

import (
	"context"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
)

func changePassword(ctx context.Context, userID, oldPassword, newPassword string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.GetUserByID(ctx, userID)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	if ok := verifyPassword(oldPassword, row.Password); !ok {
		return types.Errors.PasswordsDoNotMatch(err)
	}

	hash, err := hashPassword(newPassword)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	err = q.UpdatePassword(ctx, repo.UpdatePasswordParams{
		UserID:         userID,
		HashedPassword: hash,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

// TOOD: delete existing code
func createPasswordResetRecord(ctx context.Context, email string) (string, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.CreatePasswordResetRecord(ctx, email)
	if err != nil {
		return "", types.Errors.InternalServerError(err)
	}

	return row.Code, types.Errors.None
}

// TODO: delete code on success, verifiy age of code
func resetPassword(ctx context.Context, code, newPassword string) types.APIError {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.GetPasswordResetByCode(ctx, code)
	if err != nil {
		return types.Errors.InvalidResetCode(err)
	}

	hash, err := hashPassword(newPassword)
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	err = q.UpdatePassword(ctx, repo.UpdatePasswordParams{
		UserID:         row.ID,
		HashedPassword: hash,
	})
	if err != nil {
		return types.Errors.InternalServerError(err)
	}

	return types.Errors.None
}

package auth

import (
	"context"
	"errors"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func getUserById(ctx context.Context, userID string) (types.UserResponse, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.GetUserByID(ctx, userID)
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

func loginUser(ctx context.Context, email, password string) (types.UserResponse, types.APIError) {
	q := db.GetQueriesFromContext(ctx)
	row, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		return types.UserResponse{}, types.Errors.InvalidEmailOrPassword(err)
	}

	if !row.EmailVerified {
		return types.UserResponse{}, types.Errors.InvalidEmailOrPassword(errors.New("email not verified"))
	}

	if !verifyPassword(password, row.Password) {
		return types.UserResponse{}, types.Errors.InvalidEmailOrPassword(errors.New("invalid email or password"))
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

func registerUser(ctx context.Context, email, username, displayname, password string) (types.UserResponse, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	_, err := q.GetUserByEmail(ctx, email)
	if err == nil {
		return types.UserResponse{}, types.Errors.EmailAlreadyInUse(errors.New("email already in use"))
	}

	_, err = q.GetUserByUsername(ctx, username)
	if err == nil {
		return types.UserResponse{}, types.Errors.UsernameAlreadyInUse(errors.New("username already in use"))
	}

	hash, err := hashPassword(password)
	if err != nil {
		return types.UserResponse{}, types.Errors.InternalServerError(err)
	}

	row, err := q.CreateUser(ctx, repo.CreateUserParams{
		Email:       email,
		Username:    username,
		Displayname: displayname,
		Color:       utils.RandomColor(),
		Password:    hash,
	})
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

func verifyToken(ctx context.Context, userID string) (types.UserResponse, types.APIError) {
	q := db.GetQueriesFromContext(ctx)

	row, err := q.GetUserByID(ctx, userID)
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

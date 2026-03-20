package auth

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

// TODO: update tests to match changes

// config

var queries *repo.Queries
var conn *sql.DB

func init() {
	conn = db.GetConnection(":memory:")
	db.ConfigDatabase(conn, "../../db/schema.sql")
	queries = repo.New(conn)
}

func configCtx() context.Context {
	ctx := context.Background()
	ctx = db.AttachQueries(ctx, queries)
	return ctx
}

func insertTestUser(t *testing.T, email, username, displayname, password string) {
	t.Helper()

	color := utils.RandomColor()
	hashedPassword, err := hashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	_, err = conn.Exec(`
		INSERT INTO profile
			(email, username, displayname, color, password)
		VALUES
			(?, ?, ?, ?, ?)`,
		email, username, displayname, color, hashedPassword)
	if err != nil {
		t.Fatalf("failed to insert test user: %v", err)
	}

	t.Cleanup(func() {
		if _, err := conn.Exec(`
			DELETE FROM profile
			WHERE email = ?`,
			email); err != nil {
			t.Fatalf("failed to cleanup test user: %v", err)
		}
	})
}

// testing

func TestLoginUser(t *testing.T) {
	ctx := configCtx()

	email := "test@test.com"
	username := "testusername"
	displayname := "testdisplayname"
	password := "testpassword"

	insertTestUser(t, email, username, displayname, password)

	t.Run("invalid email", func(t *testing.T) {
		_, apiErr := loginUser(ctx, "invalid@test.com", password)

		if apiErr == types.Errors.None {
			t.Errorf("expected error, got nil")
		}

		if apiErr.Code != types.Errors.InvalidEmailOrPassword(nil).Code {
			t.Errorf("expected InvalidEmailOrPassword, got %v", apiErr.WrappedError.Error())
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		_, apiErr := loginUser(ctx, email, "wrongpassword")

		if apiErr == types.Errors.None {
			t.Errorf("expected error, got nil")
		}

		if apiErr.Code != types.Errors.InvalidEmailOrPassword(nil).Code {
			t.Errorf("expected InvalidEmailOrPassword, got %v", apiErr.WrappedError.Error())
		}
	})

	t.Run("success", func(t *testing.T) {
		user, apiErr := loginUser(ctx, email, password)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError)
		}

		if user.Email != email {
			t.Errorf("expected %v, got %v", email, user.Email)
		}

		if user.Username != username {
			t.Errorf("expected %v, got %v", username, user.Username)
		}

		if user.Displayname != displayname {
			t.Errorf("expected %v, got %v", displayname, user.Displayname)
		}
	})
}

func TestRegisterUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := configCtx()

		email := "register_success@test.com"
		username := "registersuccess"
		displayname := "register display"
		password := "testpassword"

		t.Cleanup(func() {
			if _, err := conn.Exec(`
				DELETE FROM profile
				WHERE email = ?`,
				email); err != nil {
				t.Fatalf("failed to cleanup profile: %v", err)
			}
		})

		user, apiErr := registerUser(ctx, email, username, displayname, password)

		if apiErr != types.Errors.None {
			t.Errorf("expected nil, got %v", apiErr.WrappedError.Error())
		}

		if user.Email != email {
			t.Errorf("expected %v, got %v", email, user.Email)
		}

		if user.Username != username {
			t.Errorf("expected %v, got %v", username, user.Username)
		}

		if user.Displayname != displayname {
			t.Errorf("expected %v, got %v", displayname, user.Displayname)
		}
	})

	t.Run("email already in use", func(t *testing.T) {
		ctx := configCtx()

		email := "register_email_in_use@test.com"
		username := "newusername"
		displayname := "new display"
		password := "testpassword"

		insertTestUser(t, email, "existingusername", "existing display", "existingpassword")

		_, apiErr := registerUser(ctx, email, username, displayname, password)

		if apiErr == types.Errors.None {
			t.Errorf("expected error, got nil")
		}

		if apiErr.Code != types.Errors.EmailAlreadyInUse(nil).Code {
			t.Errorf("expected EmailAlreadyInUse, got %v", apiErr.WrappedError.Error())
		}
	})

	t.Run("username already in use", func(t *testing.T) {
		ctx := configCtx()

		email := "another@test.com"
		username := "register_username_in_use"
		displayname := "new display"
		password := "testpassword"

		insertTestUser(t, "existing@test.com", username, "existing display", "existingpassword")

		_, apiErr := registerUser(ctx, email, username, displayname, password)

		if apiErr == types.Errors.None {
			t.Errorf("expected error, got nil")
		}

		if apiErr.Code != types.Errors.UsernameAlreadyInUse(nil).Code {
			t.Errorf("expected UsernameAlreadyInUse, got %v", apiErr.WrappedError.Error())
		}
	})
}

func TestRefreshToken(t *testing.T) {
	ctx := configCtx()

	email := "test@test.com"
	username := "testusername"
	displayname := "testdisplayname"
	password := "testpassword"

	// setup
	user, apiErr := registerUser(ctx, email, username, displayname, password)
	if apiErr != types.Errors.None {
		t.Fatalf("failed to register user: %v", apiErr)
	}

	// cleanup
	t.Cleanup(func() {
		if _, err := conn.Exec(`
			DELETE FROM profile
			WHERE email = ?`,
			email); err != nil {
			t.Fatalf("failed to cleanup profile: %v", err)
		}
	})

	// test
	user, apiErr = verifyToken(ctx, user.ID)

	if apiErr != types.Errors.None {
		t.Errorf("expected nil, got %v", apiErr)
	}

	if user.Email != email {
		t.Errorf("expected %v, got %v", email, user.Email)
	}

	if user.Username != username {
		t.Errorf("expected %v, got %v", username, user.Username)
	}

	if user.Displayname != displayname {
		t.Errorf("expected %v, got %v", displayname, user.Displayname)
	}
}

func TestRefreshTokenUserNotFound(t *testing.T) {
	ctx := configCtx()

	_, apiErr := verifyToken(ctx, "non-existent-user-id")

	if apiErr == types.Errors.None {
		t.Errorf("expected error, got nil")
	}

	if apiErr.Code != types.Errors.InternalServerError(nil).Code {
		t.Errorf("expected InternalServerError, got %v", apiErr.WrappedError.Error())
	}
}

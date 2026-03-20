package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/db"
)

type contextKey string

const userIDKey contextKey = "userId"

// public

func GetUserIDFromContext(ctx context.Context) string {
	return ctx.Value(userIDKey).(string)
}

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := db.GetQueriesFromContext(r.Context())
		tokenStr := getTokenFromRequest(r)
		token, err := verifyJWT(tokenStr)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userId"].(string)

		row, err := q.GetUserByID(r.Context(), userID)
		if err != nil {
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, userIDKey, row.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

// internal

func createJWT(secret []byte, userID string) (string, error) {
	expiration := time.Second * time.Duration(config.Env.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func setJWTCookie(w http.ResponseWriter, token string) {
	tokenCookie := &http.Cookie{
		Name:     config.Env.JWTKey,
		Value:    token,
		Path:     "/",
		MaxAge:   config.Env.JWTExpirationInSeconds,
		HttpOnly: true, // Don't expose to JavaScript
		Secure:   true, // Only send over HTTPS
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, tokenCookie)
}

func deleteJWTCookie(w http.ResponseWriter) {
	tokenCookie := &http.Cookie{
		Name:     config.Env.JWTKey,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true, // Don't expose to JavaScript
		Secure:   true, // Only send over HTTPS
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, tokenCookie)
}

// helpers

func getTokenFromRequest(r *http.Request) string {
	cookie, err := r.Cookie(config.Env.JWTKey)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func verifyJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})
}

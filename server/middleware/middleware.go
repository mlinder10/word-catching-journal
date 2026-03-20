package middleware

import (
	"fmt"
	"net/http"

	"github.com/mlinder10/wcj/db"
	"github.com/mlinder10/wcj/db/repo"
)

func RegisterQueries(next http.Handler, q *repo.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = db.AttachQueries(ctx, q)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer fmt.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the origin to the request's origin instead of "*"
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

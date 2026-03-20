package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlinder10/wcj/db/repo"
	"github.com/mlinder10/wcj/middleware"
	"github.com/mlinder10/wcj/service/auth"
	"github.com/mlinder10/wcj/service/define"
	"github.com/mlinder10/wcj/service/feed"
	"github.com/mlinder10/wcj/service/post"
	"github.com/mlinder10/wcj/service/profile"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()
	v1router := apirouter.PathPrefix("/v1").Subrouter()

	router.Use(middleware.Logging)
	queries := repo.New(s.db)
	router.Use(func(next http.Handler) http.Handler {
		return middleware.RegisterQueries(next, queries)
	})

	// auth
	v1router.HandleFunc("/login", auth.POST_Login).Methods("POST")
	v1router.HandleFunc("/register", auth.POST_Register).Methods("POST")
	v1router.HandleFunc("/logout", auth.POST_Logout).Methods("POST")
	v1router.HandleFunc("/verify", auth.POST_Verify).Methods("POST")
	v1router.HandleFunc("/verify/resend/{id}", auth.POST_Verify).Methods("POST")
	v1router.HandleFunc("/refresh", auth.WithJWTAuth(auth.GET_Refresh)).Methods("GET")
	v1router.HandleFunc("/reset-password", auth.WithJWTAuth(auth.PATCH_ResetPassword)).Methods("PATCH")
	v1router.HandleFunc("/reset-password", auth.WithJWTAuth(auth.POST_ResetPassword)).Methods("POST")
	v1router.HandleFunc("/reset-password/{code}", auth.WithJWTAuth(auth.POST_ResetPasswordCode)).Methods("POST")

	// feed
	v1router.HandleFunc("/feed/recent", auth.WithJWTAuth(feed.GET_FeedRecent)).Methods("GET")
	v1router.HandleFunc("/feed/following", auth.WithJWTAuth(feed.GET_FeedFollowing)).Methods("GET")

	// define
	v1router.HandleFunc("/define/{word}", auth.WithJWTAuth(define.GET_DefineWord)).Methods("GET")

	// post
	v1router.HandleFunc("/post", auth.WithJWTAuth(post.POST_Post)).Methods("POST")
	v1router.HandleFunc("/post/{postID}", auth.WithJWTAuth(post.DELETE_PostID)).Methods("GET")
	v1router.HandleFunc("/post/{postID}/like", auth.WithJWTAuth(post.POST_PostIDLike)).Methods("POST")
	v1router.HandleFunc("/post/{postID}/like", auth.WithJWTAuth(post.DELETE_PostIDLike)).Methods("DELETE")
	v1router.HandleFunc("/post/{postID}/bookmark", auth.WithJWTAuth(post.POST_PostIDBookmark)).Methods("POST")
	v1router.HandleFunc("/post/{postID}/bookmark", auth.WithJWTAuth(post.DELETE_PostIDBookmark)).Methods("DELETE")

	// profile
	v1router.HandleFunc("/profile/{profileID}", auth.WithJWTAuth(profile.GET_ProfileID)).Methods("GET")
	v1router.HandleFunc("/profile/{profileID}/posts", auth.WithJWTAuth(profile.GET_ProfileIDPosts)).Methods("GET")
	v1router.HandleFunc("/profile/{profileID}/follow", auth.WithJWTAuth(profile.POST_ProfileIDFollow)).Methods("POST")
	v1router.HandleFunc("/profile/{profileID}/follow", auth.WithJWTAuth(profile.DELETE_ProfileIDFollow)).Methods("DELETE")
	v1router.HandleFunc("/profile/{profileID}/followers", auth.WithJWTAuth(profile.GET_ProfileIDFollowers)).Methods("GET")
	v1router.HandleFunc("/profile/{profileID}/following", auth.WithJWTAuth(profile.GET_ProfileIDFollowing)).Methods("GET")

	handler := middleware.CORS(router)

	log.Println("Listening on " + s.addr)
	return http.ListenAndServe(s.addr, handler)
}

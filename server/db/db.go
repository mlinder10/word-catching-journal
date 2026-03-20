package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/mlinder10/wcj/db/repo"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func GetConnection(url string) *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite", url)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("failed to connect to db at " + url)
	}

	return db
}

type contextKey string

const queriesKey contextKey = "queries"

func AttachQueries(ctx context.Context, queries *repo.Queries) context.Context {
	return context.WithValue(ctx, queriesKey, queries)
}

func GetQueriesFromContext(ctx context.Context) *repo.Queries {
	return ctx.Value(queriesKey).(*repo.Queries)
}

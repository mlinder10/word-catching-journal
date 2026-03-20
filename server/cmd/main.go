package main

import (
	"github.com/mlinder10/wcj/cmd/api"
	"github.com/mlinder10/wcj/config"
	"github.com/mlinder10/wcj/db"
)

func main() {
	conn := db.GetConnection(config.Env.DatabaseURL)
	server := api.NewServer(config.Env.Port, conn)
	server.Run()
}

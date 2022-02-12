package main

import (
	"database/sql"
	"github.com/AI1411/golang-postgres-k8s/api"
	db "github.com/AI1411/golang-postgres-k8s/db/sqlc"
	"github.com/AI1411/golang-postgres-k8s/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		return
	}

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

package main

import (
	"database/sql"
	"github.com/AI1411/golang-postgres-k8s/api"
	db "github.com/AI1411/golang-postgres-k8s/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:15432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

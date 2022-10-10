package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/malet-pr/go-simplebank/api"
	db "github.com/malet-pr/go-simplebank/db/sqlc"
)

const (
	DBDriver          = "postgres"
	DBSource          = "postgresql://root:admin@localhost:5432/simpleBank?sslmode=disable"
	HTTPServerAddress = "0.0.0.0:8090"
)

func main() {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}

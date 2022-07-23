package main

import (
	"database/sql"
	"log"

	"github.com/Feruz666/eco-monitoring/api"
	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/Feruz666/eco-monitoring/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}

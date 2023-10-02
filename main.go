package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.comarodrigowsimple_bank/api"
	db "github.comarodrigowsimple_bank/db/sqlc"
	"github.comarodrigowsimple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config file. ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to database: ", err, config)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

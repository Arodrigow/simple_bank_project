package main

import (
	"database/sql"
	"log"

	"github.com/Arodrigow/simple_bank_project/api"
	db "github.com/Arodrigow/simple_bank_project/db/sqlc"
	"github.com/Arodrigow/simple_bank_project/util"
	_ "github.com/lib/pq"
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

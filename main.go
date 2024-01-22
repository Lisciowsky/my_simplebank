package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lisciowsky/my_simplebank/api"
	db "github.com/Lisciowsky/my_simplebank/db/sqlc"
	"github.com/Lisciowsky/my_simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	fmt.Println(config)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

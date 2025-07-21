package main

import (
	"log"
	"page-speed-server/api"
	"page-speed-server/db"
)

func main() {
	port := ":8000"

	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"page-speed-server/api"
	"page-speed-server/db"
)

func main() {
	port := os.Getenv("PORT")

	db, err := db.NewSqlStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := api.NewAPIServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
	// if len(os.Args) > 1 {
	// 	command := os.Args[1]
	// 	switch command {
	// 	case "seed":
	// 		if err := pets.SeedPets(db); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		if err := users.SeedUsers(db); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		return
	// 	case "run":
	// 		server := api.NewAPIServer(port, db)

	// 		if err := server.Run(); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		return
	// 	default:
	// 		fmt.Printf("Unknown command: %s\n", command)
	// 		fmt.Println("Usage: go run . [seed <num>] | [clean] | [serve]")
	// 		os.Exit(1)
	// 	}
	// }

}

package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"page-speed-server/services/pets"
	"page-speed-server/services/users"

	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userStore := users.NewUserStore(s.db)
	petStore := pets.NewPetStore(s.db)

	userHandler := users.NewHandler(*userStore, *petStore)
	petHander := pets.NewHandler(*petStore)

	userHandler.RegisterRoutes(router)
	petHander.RegisterRoutes(router)
	frontUrl := os.Getenv("FRONT_URL")
	frontUrlWWW := os.Getenv("FRONT_URL_WWW")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{frontUrl, frontUrlWWW},
		AllowCredentials: true,
		AllowedMethods:   []string{"OPTIONS", http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
	})
	handler := c.Handler(router)
	v1 := http.NewServeMux()
	v1.Handle("/api/v2/", http.StripPrefix("/api/v2", handler))
	server := http.Server{
		Addr:    s.addr,
		Handler: v1,
	}

	fmt.Println("listening on: ", s.addr)
	log.Fatal(server.ListenAndServe())
	return nil
}

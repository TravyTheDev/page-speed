package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"page-speed-server/services/pets"
	"page-speed-server/services/users"
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

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	userStore := users.NewUserStore(s.db)
	petStore := pets.NewPetStore(s.db)

	userHandler := users.NewHandler(*userStore, *petStore)
	petHander := pets.NewHandler(*petStore)

	userHandler.RegisterRoutes(router)
	petHander.RegisterRoutes(router)

	fmt.Println("listening on: ", s.addr)
	return server.ListenAndServe()
}

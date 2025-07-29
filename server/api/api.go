package api

import (
	"database/sql"
	"fmt"
	"net/http"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"},
		AllowCredentials: true,
		AllowedMethods:   []string{"OPTIONS", http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
	})
	handler := c.Handler(router)
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", handler))
	server := http.Server{
		Addr:    s.addr,
		Handler: v1,
	}

	fmt.Println("listening on: ", s.addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

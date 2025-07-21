package api

import (
	"database/sql"
	"fmt"
	"net/http"
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

	fmt.Println("listening on: ", s.addr)
	return server.ListenAndServe()
}

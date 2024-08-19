package api

import (
	"database/sql"
	"go-api/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address string
	db      *sql.DB
}

func NewServer(address string, db *sql.DB) *Server {
	return &Server{
		address: address,
		db:      db,
	}
}

func (server *Server) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(server.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Server running on port ", server.address)
	return http.ListenAndServe(server.address, router)
}

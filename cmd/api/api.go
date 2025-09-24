package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/TonyCross23/goecom/service/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	Addr string
	DB   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		Addr: addr,
		DB:   db,
	}
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := users.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.Addr)

	return http.ListenAndServe(s.Addr, router)
}

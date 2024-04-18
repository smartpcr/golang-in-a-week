package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"webapi/models"
	"webapi/services"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// register the handlers
	userService := services.CreateService[models.User](s.db)
	userService.RegisterRoutes(subRouter)
	projectService := services.CreateService[models.Project](s.db)
	projectService.RegisterRoutes(subRouter)
	tasksService := services.CreateService[models.Task](s.db)
	tasksService.RegisterRoutes(subRouter)

	err := http.ListenAndServe(s.address, subRouter)
	if err != nil {
		log.Fatal(err)
	}
}

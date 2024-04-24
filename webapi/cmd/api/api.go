package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"webapi/pkg/services"
	"webapi/pkg/store"
	"webapi/types"
)

type APIServer struct {
	address string
	db      *store.DbStorage
}

func NewAPIServer(address string, db *store.DbStorage) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// register the handlers
	userService := services.CreateService[types.User](s.db)
	userService.RegisterRoutes(subRouter)
	projectService := services.CreateService[types.Project](s.db)
	projectService.RegisterRoutes(subRouter)
	tasksService := services.CreateService[types.Task](s.db)
	tasksService.RegisterRoutes(subRouter)

	err := http.ListenAndServe(s.address, subRouter)
	if err != nil {
		log.Fatal(err)
	}
}

package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"webapi/pkg/inject"
	"webapi/pkg/services"
	"webapi/pkg/store"
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

func (s *APIServer) Serve(container *dig.Container) {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// register the handlers
	userService := inject.Get[*services.UserService](container)
	userService.RegisterRoutes(subRouter)
	projectService := inject.Get[*services.ProjectService](container)
	projectService.RegisterRoutes(subRouter)
	tasksService := inject.Get[*services.TaskService](container)
	tasksService.RegisterRoutes(subRouter)

	server := &http.Server{
		Addr:    s.address,
		Handler: subRouter,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-stop
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

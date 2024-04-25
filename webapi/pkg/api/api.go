package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"webapi/pkg/services"
	"webapi/pkg/store"
	"webapi/schema/v1"
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
	userService := services.CreateService[v1.User](s.db)
	userService.RegisterRoutes(subRouter)
	projectService := services.CreateService[v1.Project](s.db)
	projectService.RegisterRoutes(subRouter)
	tasksService := services.CreateService[v1.Task](s.db)
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

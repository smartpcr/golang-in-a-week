package api

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v3"
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
	app := fiber.New(fiber.Config{
		AppName: "Tasks App v1.0.0",
	})

	// register the handlers
	userService := inject.Get[*services.UserService](container)
	userService.RegisterRoutes(app)
	projectService := inject.Get[*services.ProjectService](container)
	projectService.RegisterRoutes(app)
	tasksService := inject.Get[*services.TaskService](container)
	tasksService.RegisterRoutes(app)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		if err := app.Listen(s.address); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop
	log.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err)
	}
}

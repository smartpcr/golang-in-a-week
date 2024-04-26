package main

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/dig"
	"webapi/pkg/api"
	"webapi/pkg/config"
	"webapi/pkg/inject"
	"webapi/pkg/store"
	v1 "webapi/schema/v1"
)

func main() {
	fmt.Println("Building a web API with Go!")

	container := inject.GetContainer()
	dbStore := inject.Get[*store.DbStorage](container)
	defer dbStore.Close()

	err := dbStore.EnsureTablesAndData()
	if err != nil {
		panic(err)
	}
	err = seedData(context.Background(), container)
	if err != nil {
		panic(err)
	}

	apiConfig := config.GetApiConfig()
	apiServer := api.NewAPIServer(fmt.Sprintf(":%d", apiConfig.Port), dbStore)
	apiServer.Serve(container)
}

func seedData(ctx context.Context, container *dig.Container) error {
	users := make([]*v1.User, 0)
	projects := make([]*v1.Project, 0)
	tasks := make([]*v1.Task, 0)

	userRepo := inject.Get[*store.UserRepository](container)
	usersCount, err := userRepo.Count(ctx)
	if err != nil {
		return err
	}
	if usersCount == 0 {
		users = v1.GenerateTestUsers(50)
		for _, user := range users {
			_, err := userRepo.Create(ctx, user)
			if err != nil {
				return err
			}
		}
	} else {
		users, err = userRepo.List(ctx)
		if err != nil {
			return err
		}
	}

	projectRepo := inject.Get[*store.ProjectRepository](container)
	projectsCount, err := projectRepo.Count(ctx)
	if err != nil {
		return err
	}
	if projectsCount == 0 {
		projects = v1.GenerateTestProjects(10, users)
		for _, project := range projects {
			_, err := projectRepo.Create(ctx, project)
			if err != nil {
				return err
			}
		}
	} else {
		projects, err = projectRepo.List(ctx)
		if err != nil {
			return err
		}
	}

	taskRepo := inject.Get[*store.TaskRepository](container)
	tasksCount, err := taskRepo.Count(ctx)
	if err != nil {
		return err
	}
	if tasksCount == 0 {
		tasks = v1.GenerateTestTasks(50, users, projects)
		for _, task := range tasks {
			_, err := taskRepo.Create(ctx, task)
			if err != nil {
				return err
			}
		}
	} else {
		tasks, err = taskRepo.List(ctx)
		if err != nil {
			return err
		}
	}

	log.Printf("Data seeded successfully, users=%d, projects=%d, tasks=%d", len(users), len(projects), len(tasks))

	return nil
}

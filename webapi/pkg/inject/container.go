package inject

import (
	"log"
	"sync"

	"go.uber.org/dig"
	"webapi/pkg/config"
	"webapi/pkg/services"
	"webapi/pkg/store"
	v1 "webapi/schema/v1"
)

var (
	container *dig.Container
	once      sync.Once
)

func GetContainer() *dig.Container {
	once.Do(func() {
		container = dig.New()

		// provide config
		err := container.Provide(func() (*config.DbConfig, error) {
			return config.GetDatabaseConfig()
		})
		if err != nil {
			panic(err)
		}

		// provide db storage
		err = container.Provide(func(cfg *config.DbConfig) (*store.DbStorage, error) {
			return store.NewDbStorage(cfg)
		})
		if err != nil {
			panic(err)
		}

		// provide ORM repositories
		dbStore := Get[*store.DbStorage](container)
		err = container.Provide(func() *store.UserRepository {
			return &store.UserRepository{
				OrmRepository: &store.OrmRepository[v1.User]{DbStore: dbStore},
			}
		})
		if err != nil {
			panic(err)
		}
		err = container.Provide(func() *store.ProjectRepository {
			return &store.ProjectRepository{
				OrmRepository: &store.OrmRepository[v1.Project]{DbStore: dbStore},
			}
		})
		if err != nil {
			panic(err)
		}
		err = container.Provide(func() *store.TaskRepository {
			return &store.TaskRepository{
				OrmRepository: &store.OrmRepository[v1.Task]{DbStore: dbStore},
			}
		})
		if err != nil {
			panic(err)
		}

		// provide services
		err = container.Provide(func() (*services.UserService, error) {
			repo := Get[*store.UserRepository](container)
			return &services.UserService{Repo: repo}, nil
		})
		if err != nil {
			panic(err)
		}
		err = container.Provide(func() (*services.ProjectService, error) {
			repo := Get[*store.ProjectRepository](container)
			return &services.ProjectService{Repo: repo}, nil
		})
		if err != nil {
			panic(err)
		}
		err = container.Provide(func() (*services.TaskService, error) {
			repo := Get[*store.TaskRepository](container)
			return &services.TaskService{Repo: repo}, nil
		})
		if err != nil {
			panic(err)
		}
	})

	return container
}

func Get[T any](container *dig.Container) T {
	var result T
	err := container.Invoke(func(t T) {
		result = t
	})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

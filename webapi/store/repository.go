package store

import (
	"database/sql"
	"reflect"
)

type Repository[T any] interface {
	List() ([]T, error)
	Get(id string) (T, error)
	Create(item T) (T, error)
	Update(item T) error
	Delete(id string) error
}

func CreateRepository[T any](db *sql.DB) Repository[T] {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	switch typeName {
	case "User":
		repo := &UserRepository{db: db}
		return any(repo).(Repository[T])
	case "Project":
		repo := &ProjectRepository{db: db}
		return any(repo).(Repository[T])
	case "Task":
		repo := &TaskRepository{db: db}
		return any(repo).(Repository[T])
	default:
		panic("Unknown type")
	}
}

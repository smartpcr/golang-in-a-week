package services

import (
	"database/sql"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"webapi/models"
	"webapi/store"
)

type Service interface {
	RegisterRoutes(r *mux.Router)
	HandleList(writer http.ResponseWriter, request *http.Request)
	HandleCreate(writer http.ResponseWriter, request *http.Request)
	HandleGet(writer http.ResponseWriter, request *http.Request)
	HandleUpdate(writer http.ResponseWriter, request *http.Request)
	HandleDelete(writer http.ResponseWriter, request *http.Request)
}

func CreateService[T any](db *sql.DB) Service {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	switch typeName {
	case "User":
		return &UserService{repo: store.CreateRepository[models.User](db)}
	case "Project":
		return &ProjectService{repo: store.CreateRepository[models.Project](db)}
	case "Task":
		return &TaskService{repo: store.CreateRepository[models.Task](db)}
	default:
		panic("Unknown type")
	}
}

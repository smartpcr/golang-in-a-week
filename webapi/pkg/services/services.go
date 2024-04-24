package services

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"webapi/pkg/store"
	"webapi/types"
)

type Service interface {
	RegisterRoutes(r *mux.Router)
	HandleList(writer http.ResponseWriter, request *http.Request)
	HandleCreate(writer http.ResponseWriter, request *http.Request)
	HandleGet(writer http.ResponseWriter, request *http.Request)
	HandleUpdate(writer http.ResponseWriter, request *http.Request)
	HandleDelete(writer http.ResponseWriter, request *http.Request)
}

func CreateService[T any](db *store.DbStorage) Service {
	typeName := reflect.TypeOf((*T)(nil)).Elem().Name()
	switch typeName {
	case "User":
		return &UserService{repo: store.CreateRepository[types.User](db)}
	case "Project":
		return &ProjectService{repo: store.CreateRepository[types.Project](db)}
	case "Task":
		return &TaskService{repo: store.CreateRepository[types.Task](db)}
	default:
		panic("Unknown type")
	}
}

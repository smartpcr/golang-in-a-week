package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"webapi/pkg/store"
	"webapi/types"
)

type TaskService struct {
	repo store.Repository[types.Task]
}

func (t *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", t.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/tasks", t.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", t.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", t.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", t.HandleDelete).Methods(http.MethodDelete)
}

func (t *TaskService) HandleList(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ Service = &TaskService{}

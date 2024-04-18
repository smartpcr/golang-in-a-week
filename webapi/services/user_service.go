package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"webapi/models"
	"webapi/store"
)

type UserService struct {
	repo store.Repository[models.User]
}

func (u *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", u.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/users", u.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", u.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", u.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", u.HandleDelete).Methods(http.MethodDelete)
}

func (u *UserService) HandleList(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ Service = &UserService{}

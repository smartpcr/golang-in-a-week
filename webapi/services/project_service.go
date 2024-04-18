package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"webapi/models"
	"webapi/store"
)

type ProjectService struct {
	repo store.Repository[models.Project]
}

func (p *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", p.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/projects", p.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/projects/{id}", p.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id}", p.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/projects/{id}", p.HandleDelete).Methods(http.MethodDelete)
}

func (p *ProjectService) HandleList(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

var _ Service = &ProjectService{}

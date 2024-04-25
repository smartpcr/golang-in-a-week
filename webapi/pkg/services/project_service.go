package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type ProjectService struct {
	repo store.Repository[v1.Project]
}

func (p *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", p.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/projects", p.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/projects/{id}", p.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id}", p.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/projects/{id}", p.HandleDelete).Methods(http.MethodDelete)
}

func (p *ProjectService) HandleList(writer http.ResponseWriter, request *http.Request) {
	projects, err := p.repo.List(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(projects)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (p *ProjectService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	var project v1.Project
	err := json.NewDecoder(request.Body).Decode(&project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	createdProject, err := p.repo.Create(request.Context(), &project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(createdProject)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (p *ProjectService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := p.repo.Get(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (p *ProjectService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var project v1.Project
	err = json.NewDecoder(request.Body).Decode(&project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	project.ID = uint(id)
	err = p.repo.Update(request.Context(), &project)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (p *ProjectService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.repo.Delete(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

var _ Service = &ProjectService{}

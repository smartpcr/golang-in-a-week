package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type TaskService struct {
	Repo store.Repository[v1.Task]
}

func (t *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", t.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/tasks", t.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", t.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", t.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", t.HandleDelete).Methods(http.MethodDelete)
}

func (t *TaskService) HandleList(writer http.ResponseWriter, request *http.Request) {
	tasks, err := t.Repo.List(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(tasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (t *TaskService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	var task v1.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := t.Repo.Create(request.Context(), &task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(createdTask)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (t *TaskService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := t.Repo.Get(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (t *TaskService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var task v1.Task
	err = json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = uint(id)
	err = t.Repo.Update(request.Context(), &task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (t *TaskService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.Repo.Delete(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

var _ Service = &TaskService{}

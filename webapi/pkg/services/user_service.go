package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"webapi/pkg/store"
	"webapi/schema/v1"
)

type UserService struct {
	Repo store.Repository[v1.User]
}

func (u *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", u.HandleList).Methods(http.MethodGet)
	r.HandleFunc("/users", u.HandleCreate).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", u.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", u.HandleUpdate).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", u.HandleDelete).Methods(http.MethodDelete)
}

func (u *UserService) HandleList(writer http.ResponseWriter, request *http.Request) {
	users, err := u.Repo.List(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (u *UserService) HandleCreate(writer http.ResponseWriter, request *http.Request) {
	var user v1.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := u.Repo.Create(request.Context(), &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(createdUser)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (u *UserService) HandleGet(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.Repo.Get(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (u *UserService) HandleUpdate(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var user v1.User
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = uint(id)
	err = u.Repo.Update(request.Context(), &user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (u UserService) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Repo.Delete(request.Context(), uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

var _ Service = &UserService{}

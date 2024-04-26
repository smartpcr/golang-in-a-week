package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	RegisterRoutes(r *mux.Router)
	HandleList(writer http.ResponseWriter, request *http.Request)
	HandleCreate(writer http.ResponseWriter, request *http.Request)
	HandleGet(writer http.ResponseWriter, request *http.Request)
	HandleUpdate(writer http.ResponseWriter, request *http.Request)
	HandleDelete(writer http.ResponseWriter, request *http.Request)
}

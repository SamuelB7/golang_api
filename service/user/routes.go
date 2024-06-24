package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.login).Methods("POST")
	router.HandleFunc("/signUp", handler.login).Methods("POST")
}

func (handler *Handler) login(response http.ResponseWriter, request *http.Request) {

}

func (handler *Handler) signUp(response http.ResponseWriter, request *http.Request) {

}

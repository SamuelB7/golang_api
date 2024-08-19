package user

import (
	"fmt"
	"go-api/service/auth"
	"go-api/types"
	"go-api/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.login).Methods("POST")
	router.HandleFunc("/register", handler.register).Methods("POST")
}

func (handler *Handler) login(response http.ResponseWriter, request *http.Request) {

}

func (handler *Handler) register(response http.ResponseWriter, request *http.Request) {
	var payload types.RegisterPayload

	if err := utils.ParseJson(request, &payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(response, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	_, err := handler.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(response, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return

	}

	err = handler.store.CreateUser(types.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
		Role:     payload.Role,
	})

	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(response, http.StatusCreated, map[string]string{"message": "user created"})
}

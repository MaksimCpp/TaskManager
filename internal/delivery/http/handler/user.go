package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/usecase/user"
)

type UserHandler struct {
	createUseCase user.CreateUserUseCase
	loginUseCase user.LoginUserUseCase
}

func NewUserHandler(
	createUseCase user.CreateUserUseCase, 
	loginUseCase user.LoginUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUseCase: createUseCase,
		loginUseCase: loginUseCase,
	}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request user.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	err = handler.createUseCase.Execute(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request user.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	userEntity, err := handler.loginUseCase.Execute(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := user.NewLoginUserOutput(userEntity.ID, userEntity.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

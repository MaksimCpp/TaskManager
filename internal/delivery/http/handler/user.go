package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/usecase/user"
)

type UserHandler struct {
	registerUseCase user.RegisterUserUseCase
	loginUseCase user.LoginUserUseCase
}

func NewUserHandler(
	registerUseCase user.RegisterUserUseCase, 
	loginUseCase user.LoginUserUseCase,
) *UserHandler {
	return &UserHandler{
		registerUseCase: registerUseCase,
		loginUseCase: loginUseCase,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request user.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	err = h.registerUseCase.Execute(r.Context(), request)
	if err != nil {
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request user.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	result, err := h.loginUseCase.Execute(r.Context(), request)
	if err != nil {
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": result.Token,
	})
}

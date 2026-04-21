package http_delivery

import (
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/delivery/http/handler"
)

func NewRouter(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("POST /users/login", userHandler.Login)

	return mux
}
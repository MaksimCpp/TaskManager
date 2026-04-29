package httpdelivery

import (
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/delivery/http/handler"
	"github.com/MaksimCpp/TaskManager/internal/delivery/http/middleware"
	jwtservice "github.com/MaksimCpp/TaskManager/internal/service/jwt"
)

func NewRouter(
	userHandler *handler.UserHandler,
	taskHandler *handler.TaskHandler,
	jwtService *jwtservice.JWTService,
) http.Handler {
	mux := http.NewServeMux()
	auth := middleware.Auth(jwtService)

	mux.HandleFunc("POST /users", userHandler.Register)
	mux.HandleFunc("POST /users/login", userHandler.Login)

	mux.Handle("POST /tasks", auth(http.HandlerFunc(taskHandler.Create)))
	mux.Handle("GET /tasks", auth(http.HandlerFunc(taskHandler.Get)))
	mux.Handle("DELETE /tasks", auth(http.HandlerFunc(taskHandler.Delete)))

	return mux
}
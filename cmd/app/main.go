package main

import (
	"fmt"
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/delivery/http/handler"
	http_delivery "github.com/MaksimCpp/TaskManager/internal/delivery/http"
	"github.com/MaksimCpp/TaskManager/internal/infrastructure/postgres"
	"github.com/MaksimCpp/TaskManager/internal/repository"
	"github.com/MaksimCpp/TaskManager/internal/usecase/user"
	"github.com/MaksimCpp/TaskManager/pkg/config"
)

// @title Task Manager API
// @version 1.0
// @description API для управления задачами
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.New()
	db, err := postgres.NewDB(cfg.DBUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	userRepository := repository.NewPostgreSQLUserRepository(db)
	createUserUseCase := user.NewPostgreSQLRegisterUserUseCase(userRepository)
	loginUserUseCase := user.NewPostgreSQLLoginUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(createUserUseCase, loginUserUseCase)
	router := http_delivery.NewRouter(userHandler)
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
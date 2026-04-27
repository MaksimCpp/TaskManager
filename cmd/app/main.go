package main

import (
	"fmt"
	"net/http"

	http_delivery "github.com/MaksimCpp/TaskManager/internal/delivery/http"
	"github.com/MaksimCpp/TaskManager/internal/delivery/http/handler"
	"github.com/MaksimCpp/TaskManager/internal/infrastructure/postgres"
	"github.com/MaksimCpp/TaskManager/internal/repository"
	jwt_service "github.com/MaksimCpp/TaskManager/internal/service/jwt"
	"github.com/MaksimCpp/TaskManager/internal/usecase/user"
	"github.com/MaksimCpp/TaskManager/pkg/config"
)

func main() {
	cfg := config.New()
	db, err := postgres.NewDB(cfg.DBUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	jwtService := jwt_service.NewJWTService(
		cfg.Secret,
	)
	userRepository := repository.NewPostgreSQLUserRepository(db)
	createUserUseCase := user.NewPostgreSQLRegisterUserUseCase(userRepository)
	loginUserUseCase := user.NewPostgreSQLLoginUserUseCase(userRepository, *jwtService)
	userHandler := handler.NewUserHandler(createUserUseCase, loginUserUseCase)
	router := http_delivery.NewRouter(userHandler)
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
package main

import (
	"fmt"
	"net/http"

	httpdelivery "github.com/MaksimCpp/TaskManager/internal/delivery/http"
	"github.com/MaksimCpp/TaskManager/internal/delivery/http/handler"
	"github.com/MaksimCpp/TaskManager/internal/infrastructure/postgres"
	"github.com/MaksimCpp/TaskManager/internal/repository"
	jwtservice "github.com/MaksimCpp/TaskManager/internal/service/jwt"
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
	jwtService := jwtservice.NewJWTService(
		cfg.Secret,
	)
	userRepository := repository.NewPostgreSQLUserRepository(db)
	createUserUseCase := user.NewPostgreSQLRegisterUserUseCase(userRepository)
	loginUserUseCase := user.NewPostgreSQLLoginUserUseCase(userRepository, *jwtService)
	userHandler := handler.NewUserHandler(createUserUseCase, loginUserUseCase)
	router := httpdelivery.NewRouter(userHandler)
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
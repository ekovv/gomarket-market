package main

import (
	"go-rest-api/config"
	"go-rest-api/internal/controller"
	"go-rest-api/internal/controller/router"
	"go-rest-api/internal/db"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
	"go-rest-api/internal/validator"
)

func main() {
	cnf := config.New()
	db := db.NewDB(cnf)
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	sessionSevice := service.NewSessionService()
	userUsecase := service.NewUserUseCase(userRepository, userValidator, &sessionSevice)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	authHandler "auth-service/handlers/auth"
	"auth-service/repository"
	"auth-service/services/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {

	var db *gorm.DB

	userRepository := repository.NewUserRepository(db)

	authService := auth.New(userRepository)

	authHandler := authHandler.New(authService)

	app := fiber.New()

	app.Post("/auth/login", authHandler.Login)

	app.Listen(":4000")

}

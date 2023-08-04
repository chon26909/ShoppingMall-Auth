package auth

import (
	"auth-service/services/auth"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	authService auth.Service
}

type Handler interface {
	Login(ctx *fiber.Ctx) error
}

func New(authService auth.Service) Handler {
	return &handler{
		authService: authService,
	}
}

package auth

import (
	"auth-service/dto"
	"auth-service/logger"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
)

func (h *handler) Login(ctx *fiber.Ctx) error {

	var body dto.LoginRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	logger.Debug("body", zapcore.Field{})

	fmt.Println(body)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0000",
		"data":    nil,
		"message": nil,
	})
}

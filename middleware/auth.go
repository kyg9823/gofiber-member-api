package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"

	"github.com/kyg9823/gofiber-member-api/config"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetConfig("ACCESS_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"result": err.Error(),
		})
}

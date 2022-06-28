package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kyg9823/gofiber-member-api/handlers"
)

func SetRouter(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	api.Get("/members", handlers.GetMembers)
	api.Get("/members/:id", handlers.GetMember)
	api.Post("/members/:id", handlers.NewMember)
	api.Put("/members/:id", handlers.PutMember)
	api.Delete("/members/:id", handlers.DeleteMember)

	api.Post("/auth", handlers.Login)
}

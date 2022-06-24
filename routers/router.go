package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kyg9823/gofiber-member-api/handler"
)

func SetRouter(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	api.Get("/members", handler.GetMembers)
	api.Get("/members/:id", handler.GetMember)
	api.Post("/members/:id", handler.NewMember)
	api.Put("/members/:id", handler.PutMember)
	api.Delete("/members/:id", handler.DeleteMember)
}

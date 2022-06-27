package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyg9823/gofiber-member-api/config"
	"github.com/kyg9823/gofiber-member-api/database"
	"github.com/kyg9823/gofiber-member-api/handlers"
	"github.com/kyg9823/gofiber-member-api/routers"
	"github.com/kyg9823/gofiber-member-api/utils"
)

func main() {
	database.ConnectDB()
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.DefaultErrorHandler,
	})

	app.Get("/healthcheck", handlers.Healthcheck)

	routers.SetRouter(app)

	profile := config.GetConfig("PROFILE")
	if profile == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

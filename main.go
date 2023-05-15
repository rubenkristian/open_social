package main

import (
	"rubenkristian/open-social/app/handlers"
	"rubenkristian/open-social/app/models"

	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func main() {
	app := fiber.New()
	models := models.BuildModel()

	api := app.Group("/api")

	handlers.BuildHandler(api, models)

	app.Listen(":4000")
}

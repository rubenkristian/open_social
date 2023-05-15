package main

import "github.com/gofiber/fiber/v2"

type Data struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := Data{
			Name:    "Kristian ruben",
			Picture: "ruben.jpg",
		}
		return c.JSON(data)
	})

	app.Listen(":4000")
}

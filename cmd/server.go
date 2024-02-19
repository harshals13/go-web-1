package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshals13/go-web-1/database"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

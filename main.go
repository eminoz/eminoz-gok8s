package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello from go:)")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	app.Listen(":" + port)
}

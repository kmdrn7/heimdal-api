package main

import (
	"github.com/gofiber/fiber/v2"
	"heimdal-api/cmd/heimdal"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		heimdal.Execute()
		return c.SendStatus(fiber.StatusOK)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
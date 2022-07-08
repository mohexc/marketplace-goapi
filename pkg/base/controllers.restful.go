package base

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Marketplace API is running...")
	})
}

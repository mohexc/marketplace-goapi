package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	routerAdmin := app.Group("/admin/users")

	routerAdmin.Post("/", func(c *fiber.Ctx) error {
		userReq := &UserRequest{}
		err := c.BodyParser(userReq)

		if err != nil {
			return err
		}

		userResp, err := CreateUser(userReq)
		if err != nil {
			return err
		}

		return c.JSON(userResp)
	})

	routerAdmin.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("List user")
	})

	routerAdmin.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Detail user")
	})

	routerAdmin.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update user")
	})

	routerAdmin.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete user")
	})
}

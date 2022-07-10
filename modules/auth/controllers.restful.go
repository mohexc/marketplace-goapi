package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Router(app *fiber.App) {

	routerAdmin := app.Group("/admin/users")

	routerAdmin.Post("/", func(c *fiber.Ctx) error {
		userReq := &CreateUserRequest{}
		err := c.BodyParser(userReq)
		if err != nil {
			return err
		}
		errors := ValidateStruct(*userReq)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		userResp, err := CreateUser(userReq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(userResp)
	})

	routerAdmin.Get("/", func(c *fiber.Ctx) error {
		userResp, err := GetUsers()
		if err != nil {
			return err
		}
		return c.JSON(userResp)
	})

	routerAdmin.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, errUUID := uuid.Parse(id)
		if errUUID != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		userResp, err := GetUserById(uuid)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.JSON(userResp)
	})

	routerAdmin.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, errUUID := uuid.Parse(id)
		if errUUID != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		userReq := &CreateUserRequest{}
		err := c.BodyParser(userReq)
		if err != nil {
			return err
		}
		errors := ValidateStruct(*userReq)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		userResp, err := UpdateUserById(uuid, userReq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return c.JSON(userResp)
	})

	routerAdmin.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, errUUID := uuid.Parse(id)
		if errUUID != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		err := DeleteUserById(uuid)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendString("Delete Sucess")
	})

}

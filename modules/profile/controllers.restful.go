package profile

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	routerAdmin := app.Group("/admin/profiles")

	routerAdmin.Post("/", func(c *fiber.Ctx) error {
		profileReq := &CreateProfileRequest{}
		err := c.BodyParser(profileReq)
		if err != nil {
			return err
		}
		errors := ValidateStruct(*profileReq)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		profileResp, err := CreateProfile(profileReq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(profileResp)
	})

}

package main

import (
	"fmt"
	"marketplace-goapi/modules/auth"
	"marketplace-goapi/modules/base"
	"marketplace-goapi/provider"

	"github.com/gofiber/fiber/v2"
)

func main() {
	startProvider()
	app := fiber.New()

	startRouter(app)

	app.Listen(":5500")
}

func startProvider() {
	fmt.Println("Starting provider")
	db := provider.ConnectPostgresql()
	auth.MigrateUser(db)
}

func startRouter(app *fiber.App) {
	base.Router(app)
	auth.Router(app)

}

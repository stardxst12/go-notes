package main

import(
	"github.com/gofiber/fiber/v2"
	"go-notes-api/utils"
	"go-notes-api/routes"
)

func main() {
	app := fiber.New()

	utils.ConnectDatabase()

	api := app.Group("/api")
	routes.SetupAuthRoutes(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Notes API!")
	})

	app.Listen(":3000")
}
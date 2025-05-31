package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-notes-api/handlers"
)

func SetupAuthRoutes(router fiber.Router) {
	router.Post("/register", handlers.Register)
}
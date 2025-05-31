package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"go-notes-api/handlers"
)

func SetupAuthRoutes(router fiber.Router, db *gorm.DB) {
	router.Post("/register", handlers.Register(db))
	router.Post("/login", handlers.Login(db))
}
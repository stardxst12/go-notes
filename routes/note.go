package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-notes-api/handlers"
	"go-notes-api/middleware"
	"gorm.io/gorm"
)

func SetupNoteRoutes(router fiber.Router, db *gorm.DB) {
	notes := router.Group("/notes", middleware.JWTProtected())
	notes.Post("/", handlers.CreateNote(db))
	notes.Get("/", handlers.GetNote(db))
	notes.Get("/:id", handlers.GetNoteByID(db))
	notes.Put("/:id", handlers.UpdateNote(db))
	notes.Delete("/:id", handlers.DeleteNote(db))
	// Add GET, PUT, DELETE here next
}
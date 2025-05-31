package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"go-notes-api/models"
)

func CreateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)

		var note models.Note
		if err := c.BodyParser(&note); err!= nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		note.UserID = userID
		if err := db.Create(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create note"})
		}

		return c.JSON(note)
	}
}

// We'll add other handlers like GetNotes, GetNoteByID, UpdateNote, DeleteNote next

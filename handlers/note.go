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

func GetNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)

		var notes []models.Note
		if err := db.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
		}

		return c.JSON(notes)
	}
}

// We'll add other handlers like GetNoteByID, UpdateNote, DeleteNote next

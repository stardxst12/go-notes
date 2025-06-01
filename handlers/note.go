package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"go-notes-api/models"
	"go-notes-api/utils"
	"strconv"
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

		//query params
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		title := c.Query("title", "")

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		var notes []models.Note
		query := db.Where("user_id = ?", userID)

		if title != "" {
			query = query.Where("title LIKE ?", "%"+title+"%")
		}

		result := query.Limit(limit).Offset(offset).Find(&notes)
		
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notes"})
		}

		return c.JSON(notes)
	}
}

func GetNoteByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		noteID := c.Params("id")

		var note models.Note
		result := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Note not found or unauthorized",
			})
		}

		return c.JSON(note)
	}
}

func UpdateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		noteID := c.Params("id")

		var note models.Note
		if err := db.First(&note, "id = ? AND user_id = ?", noteID, userID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}

		type UpdateNoteInput struct {
			Title string `json:"title"`
			Content string `json:"content"`
		}

		var input UpdateNoteInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		note.Title = input.Title
		note.Content = input.Content

		if err := db.Save(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update note"})
		}

		return c.JSON(note)
	}
}

func DeleteNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uint)
		noteID := c.Params("id")

		var note models.Note
		result := utils.DB.Where("id = ? AND user_id =?", noteID, userID).First(&note)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found or unauthorized"})
		}

		if err := db.Delete(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
		}
		return c.JSON(fiber.Map{"message": "Note deleted successfully"})

	}
}



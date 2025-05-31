package handlers

import (
	"go-notes-api/models"
	"go-notes-api/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		if data["name"] == "" || data["email"] == "" || data["password"] == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}

		user := models.User{
			Name:     data["name"],
			Email:    data["email"],
			Password: string(hashedPassword),
		}

		result := db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
		}

		return c.JSON(fiber.Map{
			"message": "User registered successfully",
		})
	}
}
	

func Login (db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginInput struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}

		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		if input.Email == "" || input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
		}

		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err!=nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		token, err := utils.GenerateJWT(user.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
		}

		return c.JSON(fiber.Map{"token": token})
	}
}


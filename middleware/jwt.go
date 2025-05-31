package middleware

import (
	"strings"
	"log"
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v4"
	"go-notes-api/utils"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("[JWT] Missing Authorization header")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		// Extract token from "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("[JWT] Extracted token:", token)

		// Validate the token and get the user ID
		userID, err := utils.ValidateJWT(token)
		if err != nil {
			log.Printf("[JWT Middleware] Token validation error: %v\n", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}
		log.Println("[JWT] Token valid. User ID:", userID)

		// Store the user ID in the context
		c.Locals("userID", userID)

		return c.Next()
	}
}

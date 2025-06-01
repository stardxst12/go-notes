package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenStr string) (uint, error) {
	//parsing token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("cannot convert claims to MapClaims")
	}

	fmt.Printf("JWT Claims: %+v\n", claims)

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id claim missing or invalid")
	}

	userID := uint(userIDFloat)
	return userID, nil
}


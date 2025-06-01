package main

import (
	"log"
	"fmt"
	"go-notes-api/models"
	"go-notes-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := utils.ConnectDatabase()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), 12)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	//dummy user
	for i := 1; i <= 10; i++ {
		user := models.User{
			Name:     fmt.Sprintf("Demo User %d", i),
			Email:    fmt.Sprintf("demo%d@example.com", i),
			Password: string(hashedPassword),
		}
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatal("Failed to create user:", result.Error)
		}

		//dummy notes for the user
		for j := 1; j <= 5; j++ {
			note := models.Note{
				Title:   fmt.Sprintf("Note %d for %s", j, user.Name),
				Content: fmt.Sprintf("This is note #%d for user %d", j, i),
				UserID:  user.ID,
			}
			db.Create(&note)
		}

		log.Printf("Seeded %s (%s)\n", user.Name, user.Email)
	}

	log.Println("All users and notes seeded successfully.")
}

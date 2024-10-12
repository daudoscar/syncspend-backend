package services

import (
	"syncspend/config"
	"syncspend/models"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(input UserInput) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// Save user to database
	result := config.DB.Create(&user)
	return result.Error
}

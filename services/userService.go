package services

import (
	"syncspend/config"
	"syncspend/models"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(input UserInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: string(hashedPassword),
	}

	// Save user to database
	result := config.DB.Create(&user)
	return result.Error
}

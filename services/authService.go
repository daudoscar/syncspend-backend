package services

import (
	"errors"
	"syncspend/dto"
	"syncspend/models"
	"syncspend/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) RegisterUser(data dto.CreateUserDTO) (dto.UserResponseDTO, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("failed to hash password")
	}

	user := models.User{
		Name:     data.Name,
		Username: data.Username,
		Password: string(hashedPassword),
	}

	if err := repositories.InsertUser(user); err != nil {
		return dto.UserResponseDTO{}, err
	}

	userResponse := dto.UserResponseDTO{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}

	return userResponse, nil
}

func (s *AuthService) AuthenticateUser(data dto.LoginCredentialsDTO) (dto.UserResponseDTO, error) {
	user, err := repositories.GetUserByUsername(data.Username)
	if err != nil {
		return dto.UserResponseDTO{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return dto.UserResponseDTO{}, nil
	}

	userResponse := dto.UserResponseDTO{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}

	return userResponse, nil
}

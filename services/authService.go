package services

import (
	"errors"
	"fmt"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/models"
	"syncspend/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) RegisterUser(data dto.CreateUserDTO) (dto.CredentialResponseDTO, error) {
	var profileImageURL string

	if data.Profile != nil {
		file, err := data.Profile.Open()
		if err != nil {
			return dto.CredentialResponseDTO{}, fmt.Errorf("failed to open profile image: %v", err)
		}
		defer file.Close()

		profileImageURL, err = helpers.UploadProfileImage(file, data.Profile, int(time.Now().Unix()))
		if err != nil {
			return dto.CredentialResponseDTO{}, fmt.Errorf("failed to upload profile image: %v", err)
		}
	} else {
		profileImageURL = "https://anggurproject.blob.core.windows.net/syncspend/profile/default.png"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.CredentialResponseDTO{}, errors.New("failed to hash password")
	}

	user := models.User{
		Name:     data.Name,
		Username: data.Username,
		Password: string(hashedPassword),
		Profile:  profileImageURL,
	}

	if err := repositories.InsertUser(&user); err != nil {
		return dto.CredentialResponseDTO{}, err
	}

	accessToken, err := helpers.GenerateJWT(user.Username)
	if err != nil {
		return dto.CredentialResponseDTO{}, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := helpers.GenerateRefreshToken(user.Username)
	if err != nil {
		return dto.CredentialResponseDTO{}, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	credentialResponse := dto.CredentialResponseDTO{
		ID:           user.ID,
		Name:         user.Name,
		Profile:      profileImageURL,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return credentialResponse, nil
}

func (s *AuthService) AuthenticateUser(data dto.LoginCredentialsDTO) (dto.CredentialResponseDTO, error) {
	user, err := repositories.GetUserByUsername(data.Username)
	if err != nil {
		return dto.CredentialResponseDTO{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return dto.CredentialResponseDTO{}, nil
	}

	userResponse := dto.CredentialResponseDTO{
		ID:       user.ID,
		Name:     user.Name,
		Profile:  user.Profile,
		Username: user.Username,
	}

	return userResponse, nil
}

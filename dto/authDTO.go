package dto

import (
	"mime/multipart"
)

type CreateUserDTO struct {
	Name     string                `form:"name" binding:"required"`
	Username string                `form:"username" binding:"required"`
	Password string                `form:"password" binding:"required"`
	Profile  *multipart.FileHeader `form:"profile"`
}

type LoginCredentialsDTO struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type CredentialResponseDTO struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Profile      string `json:"profile"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

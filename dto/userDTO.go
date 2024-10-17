package dto

import (
	"mime/multipart"
)

type UpdateNameDTO struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateUserDTO struct {
	ID       string                `json:"id" binding:"required"`
	Name     string                `json:"name"`
	Profile  *multipart.FileHeader `form:"profile"`
	Password string                `form:"password"`
}

type UserResponseDTO struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

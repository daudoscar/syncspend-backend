package dto

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
}

type UserResponseDTO struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

package dto

import (
	"time"
)

type CreatePlanDTO struct {
	ID_Owner    string `form:"id_owner"`
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
}

type UpdatePlanDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ID_Owner    string `json:"-"`
}

type PlanResponseDTO struct {
	ID                uint64    `json:"id"`
	ID_Owner          string    `json:"id_owner"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	InviteCode        string    `json:"invite_code,omitempty"`
	InviteCodeExpires time.Time `json:"invite_code_expires,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

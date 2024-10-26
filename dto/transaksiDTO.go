package dto

import (
	"time"
)

type CreateTransaksiDTO struct {
	Title       string  `json:"title" binding:"required"`
	ID_Payer    uint64  `json:"id_payer"`
	ID_Receiver *uint64 `json:"id_receiver,omitempty"`
	ID_Plan     uint64  `json:"id_plan" binding:"required"`
	Nominal     float64 `json:"nominal" binding:"required"`
	IsResolved  bool    `json:"is_resolved"`
}

type TransaksiResponseDTO struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	ID_Payer    uint64    `json:"id_payer"`
	ID_Receiver uint64    `json:"id_receiver,omitempty"`
	ID_Plan     uint64    `json:"id_plan,omitempty"`
	Nominal     float64   `json:"nominal"`
	IsResolved  bool      `json:"is_resolved"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

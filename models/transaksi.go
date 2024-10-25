package models

import (
	"time"
)

type Transaksi struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(255);not null"`
	ID_Payer    uint64    `gorm:"not null"`
	ID_Receiver uint64    `gorm:""`
	ID_Plan     uint64    `gorm:""`
	Nominal     int       `gorm:"not null"`
	IsResolved  bool      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Payer    User  `gorm:"foreignKey:ID_Payer;references:ID"`
	Receiver *User `gorm:"foreignKey:ID_Receiver;references:ID"`
	Plan     *Plan `gorm:"foreignKey:ID_Plan;references:ID"`
}

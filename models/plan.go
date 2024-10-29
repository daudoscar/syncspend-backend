package models

import (
	"time"
)

type Plan struct {
	ID                uint64    `gorm:"primaryKey;autoIncrement"`
	ID_Owner          uint64    `gorm:"not null"`
	Title             string    `gorm:"type:varchar(255);not null"`
	Description       string    `gorm:"type:varchar(255);not null"`
	InviteCode        string    `gorm:"type:varchar(6)"`
	InviteCodeExpires time.Time `gorm:"type:datetime"`
	IsDeleted         bool      `gorm:"default:false"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`

	Owner      User         `gorm:"foreignKey:ID_Owner;references:ID"`
	Members    []PlanMember `gorm:"foreignKey:ID_Plan;references:ID"`
	Transaksis []Transaksi  `gorm:"foreignKey:ID_Plan;references:ID"`
}

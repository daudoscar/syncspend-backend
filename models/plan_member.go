package models

import (
	"time"
)

type PlanMember struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	ID_Plan   uint64     `gorm:"not null"`
	ID_User   uint64     `gorm:"not null"`
	IsAdmin   bool       `gorm:"default:false"`
	IsDeleted bool       `gorm:"default:false"`
	JoinedAt  time.Time  `gorm:"autoCreateTime"`
	LeftAt    *time.Time `gorm:""`

	Plan *Plan `gorm:"foreignKey:ID_Plan;references:ID"`
	User *User `gorm:"foreignKey:ID_User;references:ID"`
}

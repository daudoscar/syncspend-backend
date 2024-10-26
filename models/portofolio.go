package models

type Portofolio struct {
	ID           uint64  `gorm:"primaryKey;autoIncrement"`
	ID_Owner     uint64  `gorm:"unique;not null"`
	Saving       uint64  `gorm:"not null"`
	TotalIncome  float64 `gorm:"not null"`
	TotalExpense float64 `gorm:"not null"`
	SpendingRate float64 `gorm:"-"`

	Payer User `gorm:"foreignKey:ID_Owner;references:ID"`
}

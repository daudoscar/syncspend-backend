package services

import (
	"errors"
	"syncspend/models"

	"gorm.io/gorm"
)

type PortfolioService struct {
	DB *gorm.DB
}

func NewPortfolioService(db *gorm.DB) *PortfolioService {
	return &PortfolioService{DB: db}
}

func (s *PortfolioService) GetPortfolioByOwnerAndID(idOwner uint64, portfolioID uint64) (*models.Portofolio, error) {
	var portfolio models.Portofolio

	if err := s.DB.Where("id_owner = ? AND id = ?", idOwner, portfolioID).First(&portfolio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("portfolio not found")
		}
		return nil, err
	}
	return &portfolio, nil
}

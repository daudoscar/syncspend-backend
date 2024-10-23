package services

import (
	"errors"
	"syncspend/dto"
	"syncspend/models"
	"syncspend/repositories"

	"gorm.io/gorm"
)

type PortfolioService struct {
	DB *gorm.DB
}

func NewPortfolioService(db *gorm.DB) *PortfolioService {
	return &PortfolioService{DB: db}
}

func (s *PortfolioService) GetPortfolioByOwnerAndID(data dto.GetPortoDTO) (*models.Portofolio, error) {
	portfolio, err := repositories.GetPortfolioByOwnerAndID(data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("portfolio not found")
		}
		return nil, err
	}

	return portfolio, nil
}

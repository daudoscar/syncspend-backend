package repositories

import (
	"errors"
	"syncspend/config"
	"syncspend/dto"
	"syncspend/models"

	"gorm.io/gorm"
)

type PortoRepository struct{}

func GetPortfolioByOwnerAndID(data dto.GetPortoDTO) (*models.Portofolio, error) {
	var portfolio models.Portofolio
	if err := config.DB.Where("id_owner = ? AND id = ?", data.ID, data.ID_Owner).First(&portfolio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("portfolio not found")
		}
		return nil, err
	}
	return &portfolio, nil
}

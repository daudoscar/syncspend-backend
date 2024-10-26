package repositories

import (
	"syncspend/config"
	"syncspend/models"
)

func InsertTransaksi(transaksi *models.Transaksi) error {
	if err := config.DB.Create(&transaksi).Error; err != nil {
		return err
	}
	return nil
}

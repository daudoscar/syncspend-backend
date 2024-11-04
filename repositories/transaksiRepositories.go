package repositories

import (
	"syncspend/config"
	"syncspend/models"
)

func GetTransaksiById(id uint64) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		return nil, err
	}
	return &transaksi, nil
}

func InsertTransaksi(transaksi *models.Transaksi) error {
	if err := config.DB.Create(&transaksi).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTranskasi(transaksi *models.Transaksi) error {
	if err := config.DB.Save(transaksi).Error; err != nil {
		return err
	}
	return nil
}

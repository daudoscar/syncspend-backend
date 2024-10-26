package services

import (
	"errors"
	"syncspend/dto"
	"syncspend/models"
	"syncspend/repositories"
)

type TransaksiService struct{}

func (s *TransaksiService) CreateTransaksi(data dto.CreateTransaksiDTO) (*dto.TransaksiResponseDTO, error) {
	newTransaksi := &models.Transaksi{
		Title:       data.Title,
		ID_Payer:    data.ID_Payer,
		ID_Receiver: data.ID_Receiver,
		ID_Plan:     data.ID_Plan,
		Nominal:     data.Nominal,
		IsResolved:  data.IsResolved,
	}

	if data.ID_Receiver != nil {
		newTransaksi.ID_Receiver = data.ID_Receiver
	} else {
		newTransaksi.ID_Receiver = nil
	}

	if err := repositories.InsertTransaksi(newTransaksi); err != nil {
		return nil, errors.New("failed to create transaction")
	}

	response := &dto.TransaksiResponseDTO{
		ID:          newTransaksi.ID,
		Title:       newTransaksi.Title,
		ID_Payer:    newTransaksi.ID_Payer,
		ID_Receiver: *newTransaksi.ID_Receiver,
		ID_Plan:     newTransaksi.ID_Plan,
		Nominal:     newTransaksi.Nominal,
		IsResolved:  newTransaksi.IsResolved,
		CreatedAt:   newTransaksi.CreatedAt,
		UpdatedAt:   newTransaksi.UpdatedAt,
	}

	return response, nil
}

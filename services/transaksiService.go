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

func (s *TransaksiService) ResolveTransaksi(userID, transaksiID uint64) (*dto.TransaksiResponseDTO, error) {
	TransaksiData, err := repositories.GetTransaksiById(transaksiID)
	if err != nil {
		return nil, errors.New("unable to retrieve transaksi at this moment")
	}

	if TransaksiData.ID_Payer != userID {
		return nil, errors.New("user does not have authorization to resolve transaksi")
	}

	TransaksiData.IsResolved = true

	err = repositories.UpdateTranskasi(TransaksiData)
	if err != nil {
		return nil, errors.New("unable to update transaksi at this moment")
	}

	response := &dto.TransaksiResponseDTO{
		ID:          TransaksiData.ID,
		Title:       TransaksiData.Title,
		ID_Payer:    TransaksiData.ID_Payer,
		ID_Receiver: *TransaksiData.ID_Receiver,
		ID_Plan:     TransaksiData.ID_Plan,
		Nominal:     TransaksiData.Nominal,
		IsResolved:  TransaksiData.IsResolved,
		CreatedAt:   TransaksiData.CreatedAt,
		UpdatedAt:   TransaksiData.UpdatedAt,
	}

	return response, nil
}

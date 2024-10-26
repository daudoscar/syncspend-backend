package controllers

import (
	"errors"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func CreateTransaksi(c *gin.Context) {
	var request *dto.CreateTransaksiDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.ID_Payer = userID.(uint64)

	transaksiResponse, err := (&services.TransaksiService{}).CreateTransaksi(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Transaction created successfully", transaksiResponse)
}

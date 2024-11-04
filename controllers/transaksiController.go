package controllers

import (
	"errors"
	"strconv"
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

func ResolveTransaction(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	transaksiID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid transaction ID"))
		return
	}

	transaksiResponse, err := (&services.TransaksiService{}).ResolveTransaksi(userID.(uint64), transaksiID)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Transaction resolved successfully", transaksiResponse)
}

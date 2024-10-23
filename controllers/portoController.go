package controllers

import (
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func GetPortoID(c *gin.Context) {
	var request *dto.GetPortoDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	if request.ID_Owner == 0 {
		helpers.ValidationErrorResponse(c, "Invalid request", "ID_Owner is required")
		return
	}

	if request.ID == 0 {
		helpers.ValidationErrorResponse(c, "Invalid request", "Portfolio ID is required")
		return
	}

	portfolio, err := (&services.PortfolioService{}).GetPortfolioByOwnerAndID(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Portfolio Found", portfolio)
}

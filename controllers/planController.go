package controllers

import (
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {
	var request *dto.CreatePlanDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	planResponse, err := (&services.PlanService{}).CreatePlan(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan created successfully", planResponse)
}

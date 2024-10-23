package controllers

import (
	"errors"
	"strconv"
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

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.ID_Owner = strconv.FormatUint(userID.(uint64), 10)

	planResponse, err := (&services.PlanService{}).CreatePlan(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan created successfully", planResponse)
}

func GetUserByPlan(c *gin.Context) {
	var request *dto.GetPlanDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}
}

func GetPlan(c *gin.Context) {
	var request *dto.GetPlanDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}
}

func UpdatePlan(c *gin.Context) {
	var request *dto.PlanResponseDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	if request.Title == "" {
		helpers.ValidationErrorResponse(c, "At least change a plan", "")
		return
	}

	planResponse, err := (&services.PlanService{}).UpdatePlan(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan updated succesfully", planResponse)
}

func DeletePlan(c *gin.Context) {
	var request *dto.GetPlanDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	if request.ID == 0 {
		helpers.ValidationErrorResponse(c, "Plan ID is required for deletion", "")
		return
	}

	if err := (&services.PlanService{}).DeletePlan(*request); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "Plan deleted successfully")
}

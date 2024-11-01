package controllers

import (
	"errors"
	"strconv"
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func GetPlanById(c *gin.Context) {
	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	planResponse, err := (&services.PlanService{}).GetPlanByID(planID)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "plan retrieved successfully", planResponse)
}

func CreatePlan(c *gin.Context) {
	var request *dto.CreatePlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.ID_Owner = userID.(uint64)

	planResponse, err := (&services.PlanService{}).CreatePlan(*request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "Plan created successfully", planResponse)
}

func UpdatePlan(c *gin.Context) {
	var request *dto.UpdatePlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.ID_Owner = userID.(uint64)

	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	planResponse, err := (&services.PlanService{}).UpdatePlan(planID, *request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "plan updated successfully", planResponse)
}

func DeletePlan(c *gin.Context) {
	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	if err := (&services.PlanService{}).DeletePlan(planID, userID.(uint64)); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "plan deleted successfully")
}

func RecoverPlan(c *gin.Context) {
	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	if err := (&services.PlanService{}).RecoverPlan(planID, userID.(uint64)); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "plan recovered successfully")
}

func JoinPlan(c *gin.Context) {
	var request *dto.JoinPlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.UserID = userID.(uint64)

	if err := (&services.PlanService{}).JoinPlan(request); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "Successfully joined the plan")
}

func LeavePlan(c *gin.Context) {
	var request *dto.LeavePlanDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	request.UserID = userID.(uint64)

	if err := (&services.PlanService{}).LeavePlan(request); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "Successfully left the plan")
}

func PromoteMemberPlan(c *gin.Context) {
	var request *dto.PromoteMemberDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	request.OwnerID = userID.(uint64)

	if err := (&services.PlanService{}).PromoteMemberPlan(planID, *request); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "Successfully promoted member")
}

func DemoteAdminPlan(c *gin.Context) {
	var request *dto.PromoteMemberDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
	}

	userID, exists := c.Get("userID")
	if !exists {
		helpers.ErrorResponse(c, errors.New("user not authenticated"))
		return
	}

	planID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, errors.New("invalid plan ID"))
		return
	}

	request.OwnerID = userID.(uint64)

	if err := (&services.PlanService{}).DemoteAdminPlan(planID, *request); err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponse(c, "Successfully demoted member")
}

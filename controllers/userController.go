package controllers

import (
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var request dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	if request.Name == "" && request.Profile == nil && request.Password == "" {
		helpers.ValidationErrorResponse(c, "At least one of Name, Profile, or Password must be provided", "")
		return
	}

	userResponse, err := (&services.UserService{}).UpdateUser(request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	helpers.SuccessResponseWithData(c, "User updated successfully", userResponse)
}

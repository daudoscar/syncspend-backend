package controllers

import (
	"syncspend/dto"
	"syncspend/helpers"
	"syncspend/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var request dto.CreateUserDTO

	if err := c.ShouldBind(&request); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userResponse, err := (&services.AuthService{}).RegisterUser(request)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	accessToken, err := helpers.GenerateJWT(userResponse.Username)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	refreshToken, err := helpers.GenerateRefreshToken(userResponse.Username)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	credentialResponse := dto.CredentialResponseDTO{
		ID:           userResponse.ID,
		Name:         userResponse.Name,
		Profile:      userResponse.Profile,
		Username:     userResponse.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	helpers.SuccessResponseWithData(c, "User registered successfully", credentialResponse)
}

func Login(c *gin.Context) {
	var loginRequest dto.LoginCredentialsDTO

	if err := c.ShouldBind(&loginRequest); err != nil {
		helpers.ValidationErrorResponse(c, "Invalid request", err.Error())
		return
	}

	userResponse, err := (&services.AuthService{}).AuthenticateUser(loginRequest)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	accessToken, err := helpers.GenerateJWT(userResponse.Username)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	refreshToken, err := helpers.GenerateRefreshToken(userResponse.Username)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	credentialResponse := dto.CredentialResponseDTO{
		ID:           userResponse.ID,
		Name:         userResponse.Name,
		Username:     userResponse.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	helpers.SuccessResponseWithData(c, "Login successful", credentialResponse)
}

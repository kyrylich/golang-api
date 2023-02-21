package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/middlewares"
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"golangpet/internal/models/reader"
	"golangpet/internal/service/auth"
	"golangpet/internal/translation"
	"golangpet/internal/validation"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	username, ok := c.Get(middlewares.GinContextCurrentUsername)
	if !ok {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", "User associated with token not found")
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	userReader := reader.NewUserReader(models.DB)
	userOutput := userReader.GetByUsername(fmt.Sprintf("%s", username))

	if userOutput == nil {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", fmt.Sprintf("User with username `%s` not found", username))
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	c.JSON(http.StatusOK, userOutput)
}

func SignIn(c *gin.Context) {
	var userInput input.SignInUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}

	var authService = auth.NewAuthService(nil, nil)

	model, errorResponse := authService.SignIn(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusOK, model)
}

func SignUp(c *gin.Context) {
	var userInput input.SignUpUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}
	var authService = auth.NewAuthService(nil, nil)

	model, errorResponse := authService.SignUp(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, model)
}

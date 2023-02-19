package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"golangpet/internal/security"
	"golangpet/internal/service/auth"
	"golangpet/internal/translation"
	"golangpet/internal/validation"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	username, ok := c.Get("CURRENT_USERNAME")
	if !ok {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", "User associated with token not found")
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	user := &models.User{Username: fmt.Sprintf("%s", username)}
	models.DB.First(user, user)

	if user == nil {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", fmt.Sprintf("User with username `%s` not found", username))
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SignIn(c *gin.Context) {
	var userInput input.SignInUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}

	passwordHasher := security.BcryptPasswordHasher{}
	var authService = auth.NewAuthService(passwordHasher)

	model, errorResponse := authService.SignIn(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusOK, model)
}

func SignUp(c *gin.Context) {
	/* TODO:
	   Provide tests
	*/

	passwordHasher := security.BcryptPasswordHasher{}
	var authService = auth.NewAuthService(passwordHasher)

	var userInput input.SignUpUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}

	model, errorResponse := authService.SignUp(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, model)
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/middleware"
	"golangpet/internal/database/reader"
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/mapper"
	"golangpet/internal/service/auth"
	"golangpet/internal/translation"
	"golangpet/internal/validation"
	"net/http"
)

type AuthHandlerInterface interface {
	GetCurrentUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type AuthHandler struct {
	userReader  reader.UserReaderInterface
	authService auth.AuthServiceInterface
}

func NewAuthHandler(userReader reader.UserReaderInterface, authService auth.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{userReader: userReader, authService: authService}
}

func (a AuthHandler) GetCurrentUser(c *gin.Context) {
	username, ok := c.Get(middleware.GinContextCurrentUsername)
	if !ok {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", "User associated with token not found")
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	userOutput := mapper.MapUserModelToUserOutput(a.userReader.GetByUsername(fmt.Sprintf("%s", username)))

	if userOutput == nil {
		errResponse := output.NewErrorResponse(http.StatusNotFound)
		errResponse.AddError("", fmt.Sprintf("User with username `%s` not found", username))
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	c.JSON(http.StatusOK, userOutput)
}

func (a AuthHandler) SignIn(c *gin.Context) {
	var userInput input.SignInUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}

	model, errorResponse := a.authService.SignIn(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusOK, model)
}

func (a AuthHandler) SignUp(c *gin.Context) {
	var userInput input.SignUpUserInput

	if validationErr := c.ShouldBindJSON(&userInput); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, validation.CreateValidationResponse(translation.Translator, validationErr))
		return
	}

	model, errorResponse := a.authService.SignUp(userInput)
	if errorResponse != nil {
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, model)
}

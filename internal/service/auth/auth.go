package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golangpet/internal/config"
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"golangpet/internal/models/writer"
	"golangpet/internal/security"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type AuthServiceInterface interface {
	SignUp(input input.SignUpUserInput) (*output.SignUpUserOutput, *output.ErrorResponse)
	SignIn(input input.SignInUserInput) (*output.SignInUserOutput, *output.ErrorResponse)
}

type AuthService struct {
	passwordHasher security.PasswordHasherInterface
	userWriter     writer.UserWriterInterface
}

func NewAuthService(passwordHasher security.PasswordHasherInterface, userWriter writer.UserWriterInterface) AuthServiceInterface {
	if passwordHasher == nil {
		passwordHasher = security.BcryptPasswordHasher{}
	}
	if userWriter == nil {
		userWriter = writer.NewUserWriter(models.DB)
	}

	return &AuthService{passwordHasher: passwordHasher, userWriter: userWriter}
}

func (a *AuthService) SignIn(input input.SignInUserInput) (*output.SignInUserOutput, *output.ErrorResponse) {
	user := &models.User{Username: input.Username}

	models.DB.First(user, user)

	if user == nil || !a.passwordHasher.Verify(input.Password, user.Password) {
		errResponse := output.NewErrorResponse(http.StatusBadRequest)
		errResponse.AddError("username", "Incorrect username or password")
		return nil, errResponse
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: input.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	signedToken, err := token.SignedString(config.GetSigningKey())
	if err != nil {
		errResponse := output.NewErrorResponse(http.StatusBadRequest)
		errResponse.AddError("username", err.Error())
		return nil, errResponse
	}

	return &output.SignInUserOutput{Token: signedToken}, nil
}

func (a *AuthService) SignUp(input input.SignUpUserInput) (*output.SignUpUserOutput, *output.ErrorResponse) {
	hashedPass, err := a.passwordHasher.Hash(input.Password)

	errorResponse := output.NewErrorResponse(http.StatusBadRequest)
	if err != nil {
		errorResponse.AddError("", err.Error())
		return nil, errorResponse
	}
	input.Password = hashedPass

	userOutput, db := a.userWriter.Create(input)

	if db.Error != nil {
		errorResponse.Code = http.StatusInternalServerError
		errorResponse.AddError("", err.Error())
		return nil, errorResponse
	}

	if db.RowsAffected == 0 {
		errorResponse.AddError("username", fmt.Sprintf("User with username `%s` already exists", input.Username))
		return nil, errorResponse
	}

	return &userOutput, nil
}

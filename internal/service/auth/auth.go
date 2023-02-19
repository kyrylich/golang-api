package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"golangpet/internal/security"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type AuthServiceInterface interface {
	SignUp(input input.SignUpUserInput) (*models.User, *output.ErrorResponse)
	SignIn(input input.SignInUserInput) (*output.SignInOutput, *output.ErrorResponse)
}

type AuthService struct {
	passwordHasher security.PasswordHasher
	// TODO: Add User Writer
}

func NewAuthService(passwordHasher security.PasswordHasher) AuthServiceInterface {
	return &AuthService{passwordHasher: passwordHasher}
}

func (a *AuthService) SignIn(input input.SignInUserInput) (*output.SignInOutput, *output.ErrorResponse) {
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

	signedToken, err := token.SignedString([]byte("ThisIsOurSecret"))
	if err != nil {
		errResponse := output.NewErrorResponse(http.StatusBadRequest)
		errResponse.AddError("username", err.Error())
		return nil, errResponse
	}

	return &output.SignInOutput{Token: signedToken}, nil
}

func (a *AuthService) SignUp(input input.SignUpUserInput) (*models.User, *output.ErrorResponse) {
	hashedPass, err := a.passwordHasher.Hash(input.Password)

	errorResponse := output.NewErrorResponse(http.StatusBadRequest)
	if err != nil {
		errorResponse.AddError("", err.Error())
		return nil, errorResponse
	}

	user := models.User{Username: input.Username, Password: hashedPass}
	db := models.DB.FirstOrCreate(&user, models.User{Username: input.Username})

	if db.Error != nil {
		errorResponse.Code = http.StatusInternalServerError
		errorResponse.AddError("", err.Error())
		return nil, errorResponse
	}

	if db.RowsAffected == 0 {
		errorResponse.AddError("username", fmt.Sprintf("User with username `%s` already exists", input.Username))
		return nil, errorResponse
	}

	return &user, nil
}

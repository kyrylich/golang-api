package handlers

import (
	"fmt"
	"golangpet/internal/dto/input"
	"golangpet/tests/util"
	"net/http"
	"testing"
)

const signUpPath = "/api/auth/sign-up"
const signInPath = "/api/auth/sign-in"
const getCurrentUserPath = "/api/auth/get-current-user"

func TestSignUpShouldRegisterUserSuccessfully(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignUpUserInput{
		Username: "test12341234124124252513525",
		Password: "SuperSecurePass1234",
	}

	// Act
	result := e.POST(signUpPath).
		WithJSON(json).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	// Assert
	result.HasValue("username", json.Username)
}

func TestSignUpWithDuplicatedUserShouldReturnBadRequest(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignUpUserInput{
		Username: "TestUser",
		Password: "SuperSecurePass1234",
	}

	util.CreateUser(nil)

	// Act
	result := e.POST(signUpPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// Assert
	result.Value("code").
		IsNumber().
		IsEqual(http.StatusBadRequest)

	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual(fmt.Sprintf("User with username `%s` already exists", json.Username))
}

func TestSignUpInputValidationShouldReturnErrors(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignUpUserInput{
		Username: "1",
		Password: "1",
	}

	// Act
	result := e.POST(signUpPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	result.Value("code").
		IsNumber().
		IsEqual(http.StatusBadRequest)

	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual("Username must be at least 5 characters in length")

	result.
		Value("errors").
		Array().
		Value(1).
		Object().
		Value("message").
		IsEqual("Password must be at least 8 characters in length")
}

func TestSignInShouldReturnAuthToken(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "SuperSecurePass1234",
	}

	util.CreateUser(nil)

	// Act
	result := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusOK).JSON().Object()

	// Assert
	result.Value("token").IsString()
}

func TestSignInWithNonExistedUserShouldReturnError(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "SuperSecurePass1234",
	}

	// Act
	result := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// Assert
	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual("Incorrect username or password")
}

func TestSignInIncorrectPasswordShouldReturnError(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "this password is not mine",
	}

	util.CreateUser(nil)

	// Act
	result := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// Assert
	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual("Incorrect username or password")
}

func TestSignInWithEmptyUsernameShouldReturnError(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "",
		Password: "SuperSecurePass1234",
	}

	// Act
	result := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// Assert
	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual("Username is a required field")
}

func TestSignInWithEmptyPasswordShouldReturnError(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "",
	}

	// Act
	result := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusBadRequest).JSON().Object()

	// Assert
	result.
		Value("errors").
		Array().
		Value(0).
		Object().
		Value("message").
		IsEqual("Password is a required field")
}

func TestGetCurrentUserShouldReturnUserAssociatedWithThisToken(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "SuperSecurePass1234",
	}

	util.CreateUser(nil)

	token := e.POST(signInPath).
		WithJSON(json).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Raw()["token"]

	// Act
	result := e.GET(getCurrentUserPath).
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	result.Value("username").IsEqual(json.Username)
}

func TestGetCurrentUserWithoutAuthHeaderShouldReturnStatusUnauthorized(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	// Act
	e.GET(getCurrentUserPath).
		Expect().
		Status(http.StatusUnauthorized)
}

func TestGetCurrentUserWithEmptyTokenShouldReturnStatusUnauthorized(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	// Act
	e.GET(getCurrentUserPath).
		WithHeader("Authorization", "Bearer ").
		Expect().
		Status(http.StatusUnauthorized)
}

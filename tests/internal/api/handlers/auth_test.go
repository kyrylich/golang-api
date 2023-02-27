package handlers

import (
	"fmt"
	"golangpet/internal/dto/input"
	"golangpet/tests/util"
	"net/http"
	"testing"
)

const path = "/api/auth/sign-up"

func TestSignUpShouldRegisterUserSuccessfully(t *testing.T) {
	// Arrange
	e, serverClose := util.SetupApi(t)
	defer serverClose()

	json := input.SignInUserInput{
		Username: "test12341234124124252513525",
		Password: "SuperSecurePass1234",
	}

	// Act
	result := e.POST(path).
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

	json := input.SignInUserInput{
		Username: "TestUser",
		Password: "SuperSecurePass1234",
	}

	util.CreateUser(nil)

	// Act
	result := e.POST(path).
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

	json := input.SignInUserInput{
		Username: "1",
		Password: "1",
	}

	// Act
	result := e.POST(path).
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

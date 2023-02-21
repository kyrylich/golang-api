package output

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorResponse(t *testing.T) {
	// Arrange
	a := assert.New(t)
	expectedCode := 101

	// Act
	errorResponse := NewErrorResponse(expectedCode)

	// Assert
	a.Equal(expectedCode, errorResponse.Code)
	a.Empty(errorResponse.Errors)
}

func TestErrorResponse_AddError(t *testing.T) {
	// Arrange
	a := assert.New(t)
	field := "Field"
	message := "Error message"
	length := 5

	// Act
	errorResponse := NewErrorResponse(101)
	for i := 0; i < length; i++ {
		errorResponse.AddError(field, message)
	}

	// Assert
	a.Len(errorResponse.Errors, length)
	a.Equal(errorResponse.Errors[0].Field, field)
	a.Equal(errorResponse.Errors[0].Message, message)
}

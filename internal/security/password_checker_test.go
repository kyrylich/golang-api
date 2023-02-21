package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcryptPasswordHasher_HashAndVerify(t *testing.T) {
	// Arrange
	a := assert.New(t)

	hasher := &BcryptPasswordHasher{}
	password := "SuperSecretPassword123456"

	// Act
	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatal(err)
	}
	passwordIsOk := hasher.Verify(password, hash)

	// Assert
	a.NotEmpty(hash)
	a.True(passwordIsOk)
}

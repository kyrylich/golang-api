package security

import "golang.org/x/crypto/bcrypt"

type PasswordHasherInterface interface {
	Hash(password string) (string, error)
	Verify(password, hash string) bool
}

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() PasswordHasherInterface {
	return &BcryptPasswordHasher{}
}

func (*BcryptPasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (*BcryptPasswordHasher) Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

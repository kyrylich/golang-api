package factory

import "golangpet/internal/security"

func (f *DependencyFactory) createBcryptPasswordHasher() security.PasswordHasherInterface {
	return security.NewBcryptPasswordHasher()
}

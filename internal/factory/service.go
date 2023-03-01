package factory

import "golangpet/internal/service/auth"

func (f *DependencyFactory) createAuthService() auth.AuthServiceInterface {
	return auth.NewAuthService(f.createBcryptPasswordHasher(), f.createUserWriter(), f.createUserReader())
}

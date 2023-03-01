package factory

import (
	"golangpet/internal/api/handler"
	"gorm.io/gorm"
)

type DependencyFactory struct {
	db *gorm.DB
}

func NewDependencyFactory(db *gorm.DB) *DependencyFactory {
	return &DependencyFactory{db: db}
}

func (f *DependencyFactory) CreateAuthHandler() handler.AuthHandlerInterface {
	return handler.NewAuthHandler(f.createUserReader(), f.createAuthService())
}

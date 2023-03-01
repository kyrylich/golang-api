package reader

import (
	"golangpet/internal/model"
	"gorm.io/gorm"
)

type UserReaderInterface interface {
	GetByUsername(username string) *model.User
}

func NewUserReader(db *gorm.DB) UserReaderInterface {
	return &UserReader{db: db}
}

type UserReader struct {
	db *gorm.DB
}

func (r *UserReader) GetByUsername(username string) *model.User {
	user := &model.User{Username: username}
	r.db.First(user, user)

	return user
}

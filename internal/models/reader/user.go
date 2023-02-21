package reader

import (
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"gorm.io/gorm"
)

type UserReaderInterface interface {
	GetByUsername(username string) *output.UserOutput
}

func NewUserReader(db *gorm.DB) UserReaderInterface {
	return &UserReader{db: db}
}

type UserReader struct {
	db *gorm.DB
}

func (r *UserReader) GetByUsername(username string) *output.UserOutput {
	user := &models.User{Username: username}
	models.DB.First(user, user)

	return &output.UserOutput{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

package writer

import (
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/models"
	"gorm.io/gorm"
)

type UserWriterInterface interface {
	Create(input input.SignUpUserInput) (output.SignUpUserOutput, *gorm.DB)
}

func NewUserWriter(db *gorm.DB) UserWriterInterface {
	return &UserWriter{db: db}
}

type UserWriter struct {
	db *gorm.DB
}

func (u *UserWriter) Create(input input.SignUpUserInput) (output.SignUpUserOutput, *gorm.DB) {
	user := models.User{Username: input.Username, Password: input.Password}

	db := models.DB.FirstOrCreate(&user, models.User{Username: input.Username})

	return output.SignUpUserOutput{Username: user.Username, CreatedAt: user.CreatedAt}, db
}

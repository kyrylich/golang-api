package writer

import (
	"golangpet/internal/dto/input"
	"golangpet/internal/dto/output"
	"golangpet/internal/model"
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
	user := model.User{Username: input.Username, Password: input.Password}

	db := u.db.FirstOrCreate(&user, model.User{Username: input.Username})

	return output.SignUpUserOutput{Username: user.Username, CreatedAt: user.CreatedAt}, db
}

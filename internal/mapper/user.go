package mapper

import (
	"golangpet/internal/dto/output"
	"golangpet/internal/model"
)

func MapUserModelToUserOutput(u *model.User) *output.UserOutput {
	return &output.UserOutput{
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

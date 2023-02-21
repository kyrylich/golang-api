package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

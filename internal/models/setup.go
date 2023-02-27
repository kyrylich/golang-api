package models

import (
	"golangpet/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *config.DatabaseConfig) error {
	database, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if err != nil {
		return err
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	DB = database

	return nil
}

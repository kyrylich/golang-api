package model

import (
	"golangpet/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase(config *config.DatabaseConfig) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	database, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	db = database

	return database, nil
}

package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "user:secret@tcp(127.0.0.1:3307)/main?charset=utf8mb4&parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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

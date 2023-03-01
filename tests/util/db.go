package util

import (
	"fmt"
	"golangpet/internal/model"
	"golangpet/internal/security"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

func CleanUpDatabase(db *gorm.DB) {
	database = db

	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	for _, table := range tables {
		if _, err = sql.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)); err != nil {
			return
		}
	}
}

func CreateUser(user *model.User) *model.User {
	hasher := security.BcryptPasswordHasher{}
	pass, _ := hasher.Hash("SuperSecurePass1234")
	if user == nil {
		user = &model.User{Username: "TestUser", Password: pass}
	}
	database.Create(user)

	return user
}

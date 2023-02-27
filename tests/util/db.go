package util

import (
	"fmt"
	"golangpet/internal/models"
	"golangpet/internal/security"
	"gorm.io/gorm"
	"log"
)

func CleanUpDatabase(db *gorm.DB) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	for _, table := range tables {
		sql.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
	}

	return
}

func CreateUser(user *models.User) *models.User {
	hasher := security.BcryptPasswordHasher{}
	pass, _ := hasher.Hash("SuperSecurePass1234")
	if user == nil {
		user = &models.User{Username: "TestUser", Password: pass}
	}
	models.DB.Create(user)

	return user
}

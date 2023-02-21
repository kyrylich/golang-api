package main

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/handlers"
	"golangpet/internal/api/middlewares"
	"golangpet/internal/config"
	"golangpet/internal/models"
	"golangpet/internal/translation"
	"log"
)

func init() {
	if err := translation.RegisterValidationTranslations(); err != nil {
		log.Fatal(err.Error())
	}

	if err := config.InitConfiguration(); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	r := gin.Default()

	r.POST("/api/auth/sign-up", handlers.SignUp)
	r.POST("/api/auth/sign-in", handlers.SignIn)
	r.GET("/api/auth/get-current-user", middlewares.EnsureAuthorized, handlers.GetCurrentUser)

	models.ConnectDatabase()

	r.Run()
}

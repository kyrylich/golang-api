package main

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/rest"
	"golangpet/internal/models"
	"golangpet/internal/service/auth"
	"golangpet/internal/translation"
	"log"
	"net/http"
	"strings"
)

func init() {
	err := translation.RegisterValidationTranslations()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/api/auth/sign-up", rest.SignUp)
	r.POST("/api/auth/sign-in", rest.SignIn)
	r.GET("/api/auth/get-current-user", EnsureAuthorized, rest.GetCurrentUser)

	models.ConnectDatabase()

	r.Run()
}

func EnsureAuthorized(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username, err := auth.ParseToken(headerParts[1], []byte("ThisIsOurSecret"))
	if err != nil {
		status := http.StatusBadRequest
		if err == auth.ERROR_INVALID_ACCESS_TOKEN {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatusJSON(status, err.Error())
		return
	}
	c.Set("CURRENT_USERNAME", username)
}

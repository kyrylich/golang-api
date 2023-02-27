package router

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/handlers"
	"golangpet/internal/api/middlewares"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/auth/sign-up", handlers.SignUp)
	r.POST("/api/auth/sign-in", handlers.SignIn)
	r.GET("/api/auth/get-current-user", middlewares.EnsureAuthorized, handlers.GetCurrentUser)

	return r
}

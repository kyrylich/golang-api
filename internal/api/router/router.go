package router

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/api/middleware"
	"golangpet/internal/factory"
)

func CreateRouter(factory *factory.DependencyFactory) *gin.Engine {
	r := gin.Default()

	authHandler := factory.CreateAuthHandler()

	r.POST("/api/auth/sign-up", authHandler.SignUp)
	r.POST("/api/auth/sign-in", authHandler.SignIn)
	r.GET("/api/auth/get-current-user", middleware.EnsureAuthorized, authHandler.GetCurrentUser)

	return r
}

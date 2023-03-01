package middleware

import (
	"github.com/gin-gonic/gin"
	"golangpet/internal/app_error"
	"golangpet/internal/config"
	"golangpet/internal/service/auth"
	"net/http"
	"strings"
)

const GinContextCurrentUsername = "CURRENT_USERNAME"

const authorizationHeader = "Authorization"
const bearerPrefix = "Bearer"

func EnsureAuthorized(c *gin.Context) {
	authHeader := c.GetHeader(authorizationHeader)
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != bearerPrefix {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username, err := auth.ParseToken(headerParts[1], config.GetSigningKey())
	if err != nil {
		status := http.StatusBadRequest
		if err == app_error.ErrorInvalidAccessToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatusJSON(status, err.Error())
		return
	}
	c.Set(GinContextCurrentUsername, username)
}

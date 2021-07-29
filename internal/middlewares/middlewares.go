package middlewares

import (
	"net/http"

	"github.com/HellstromIT/authservice/internal/config"
	"github.com/HellstromIT/authservice/pkg/auth"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request, &jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtSecretMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("JWT_SECRET", jwtSecret)
		c.Next()
	}
}

func ConfigMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Config", config)
		db := config.Initialize()
		c.Set("DB", db)
		c.Next()
	}
}

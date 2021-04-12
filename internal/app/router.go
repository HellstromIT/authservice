package app

import (
	"github.com/HellstromIT/auth/cmd/authservice/internal/config"
	"github.com/HellstromIT/auth/cmd/authservice/internal/controllers"
	"github.com/HellstromIT/auth/cmd/authservice/internal/middlewares"
)

func route(c *config.Config) {
	router.GET("/user", middlewares.TokenAuthMiddleware(c.JWT.Secret), middlewares.JwtSecretMiddleware(c.JWT.Secret), controllers.CreateUser)
	router.POST("/login", middlewares.JwtSecretMiddleware(c.JWT.Secret), controllers.Login)
	router.POST("/logout", middlewares.TokenAuthMiddleware(c.JWT.Secret), middlewares.JwtSecretMiddleware(c.JWT.Secret), middlewares.ConfigMiddleware(c), controllers.LogOut)
}

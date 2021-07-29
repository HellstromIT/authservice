package controllers

import (
	"net/http"

	"github.com/HellstromIT/authservice/internal/models"
	"github.com/HellstromIT/authservice/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	jwtSecret := c.Keys["JWT_SECRET"].(string)
	if jwtSecret == "" {
		c.JSON(http.StatusInternalServerError, "Please try again")
		return
	}
	au, err := auth.ExtractTokenAuth(c.Request, &jwtSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	if !models.Model.IsAdmin(au) {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	user, err := models.Model.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

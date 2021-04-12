package controllers

import (
	"fmt"
	"net/http"

	"github.com/HellstromIT/auth/cmd/authservice/internal/models"
	"github.com/HellstromIT/auth/cmd/authservice/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	jwt_secret := c.Keys["JWT_SECRET"].(string)
	if jwt_secret == "" {
		c.JSON(http.StatusInternalServerError, "Please try again")
		return
	}
	au, err := auth.ExtractTokenAuth(c.Request, &jwt_secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	fmt.Println(au.Admin)
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

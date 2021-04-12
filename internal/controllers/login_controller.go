package controllers

import (
	"log"
	"net/http"

	"github.com/HellstromIT/authservice/internal/crypt"
	"github.com/HellstromIT/authservice/internal/models"
	"github.com/HellstromIT/authservice/internal/service"
	"github.com/HellstromIT/authservice/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	jwt_secret := c.Keys["JWT_SECRET"].(string)
	if jwt_secret == "" {
		c.JSON(http.StatusInternalServerError, "Please try again")
	}
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// check if user exists
	user, err := models.Model.GetUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// Verify Password
	verified, err := crypt.ComparePasswords([]byte(user.Password), []byte(u.Password))
	if !verified {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	// If the user is logged out the auth token is destroyed so we need to recreate it
	authData, err := models.Model.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID
	authD.Admin = user.Admin

	token, loginErr := service.Authorize.SignIn(authD, &jwt_secret)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try to login again")
		return
	}

	c.JSON(http.StatusOK, token)

}

func LogOut(c *gin.Context) {
	jwt_secret := c.Keys["JWT_SECRET"].(string)
	if jwt_secret == "" {
		c.JSON(http.StatusInternalServerError, "Please try again")
	}
	au, err := auth.ExtractTokenAuth(c.Request, &jwt_secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	delErr := models.Model.DeleteAuth(au)
	if delErr != nil {
		log.Println(delErr)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

package app

import (
	"log"

	"github.com/HellstromIT/authservice/internal/config"
	"github.com/HellstromIT/authservice/internal/models"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	var config config.Config
// 	config.Read("../../config.yaml")
// 	config.ReadEnv()
// }

var router = gin.Default()

func InitUser(email string, password string, admin int64) {
	var user = models.User{}
	user.Email = email
	user.Password = password
	user.Admin = int64(admin)

	models.Model.CreateUser(&user)
}

func StartApp() {
	var config config.Config
	config.Read("../../config.yaml")
	config.ReadEnv()

	DB := config.Initialize()

	models.Model.AutoMigrate(DB)

	route(&config)

	log.Fatal(router.Run(":" + "9001"))
}

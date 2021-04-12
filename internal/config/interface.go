package config

import (
	"gorm.io/gorm"
)

type Configurer interface {
	Initialize() *gorm.DB
	Read(string)
	ReadEnv()
}

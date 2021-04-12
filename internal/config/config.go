package config

import (
	"log"
	"os"

	"github.com/HellstromIT/authservice/internal/database"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Config struct {
	DB  Database `yaml:"database"`
	JWT Jwt      `yaml:"jwt"`
}

type Database struct {
	Type     string `yaml:"type" envconfig:"DATABASE_TYPE"`
	Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
	Port     string `yaml:"port" envconfig:"DATABASE_PORT"`
	DBName   string `yaml:"database" envconfig:"DATABASE_DATABASE"`
	Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
	User     string `yaml:"user" envconfig:"DATABASE_USER"`
	SslMode  string `yaml:"sslmode" envconfig:"DATABASE_SSLMODE"`
	TimeZone string `yaml:"timezone" envconfig:"DATABASE_TIMEZONE"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

func (c *Config) Initialize() *gorm.DB {

	var dbClient *gorm.DB
	switch {
	case c.DB.Type == "sqlite":
		var db database.Sqlite
		dbClient = db.Connect(
			c.DB.Host,
			c.DB.Port,
			c.DB.DBName,
			c.DB.Password,
			c.DB.User,
			c.DB.SslMode,
			c.DB.TimeZone)
	case c.DB.Type == "postgres":
		var db database.Postgres
		dbClient = db.Connect(
			c.DB.Host,
			c.DB.Port,
			c.DB.DBName,
			c.DB.Password,
			c.DB.User,
			c.DB.SslMode,
			c.DB.TimeZone)
	}

	return dbClient
}

func (c *Config) Read(f string) {
	buf, err := os.Open(f)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer buf.Close()

	decoder := yaml.NewDecoder(buf)
	err = decoder.Decode(c)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (c *Config) ReadEnv() {
	err := envconfig.Process("", c)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

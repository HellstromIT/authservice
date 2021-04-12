package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
}

func (p *Postgres) Connect(Host string, Port string, DBName string, Password string, User string, SslMode string, TimeZone string) *gorm.DB {
	dsn := ("host=" + Host + " user=" + User + " password=" + Password + " dbname=" + DBName + " port=" + Port + " sslmode=" + SslMode + " TimeZone=" + TimeZone)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return database
}

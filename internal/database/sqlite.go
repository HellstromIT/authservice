package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
}

func (p *Sqlite) Connect(Host string, Port string, DBName string, Password string, User string, SslMode string, TimeZone string) *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return database
}

package models

import (
	"github.com/HellstromIT/auth/cmd/authservice/pkg/auth"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

var (
	Model modelInterface = &Server{}
)

type modelInterface interface {
	Initialize(*gorm.DB)
	AutoMigrate(*gorm.DB)

	// user methods
	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

	// auth methods
	FetchAuth(*auth.AuthDetails) (*Auth, error)
	DeleteAuth(*auth.AuthDetails) error
	CreateAuth(uint64) (*Auth, error)
	IsAdmin(*auth.AuthDetails) bool
}

func (s *Server) Initialize(db *gorm.DB) {
	s.DB = db
}

func (s *Server) AutoMigrate(db *gorm.DB) {
	s.DB = db

	s.DB.Debug().AutoMigrate(
		&User{},
		&Auth{},
	)
}

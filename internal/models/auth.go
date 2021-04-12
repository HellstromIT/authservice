package models

import (
	"fmt"

	"github.com/HellstromIT/authservice/pkg/auth"
	"github.com/gofrs/uuid"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
	Admin    int64  `gorm:"default:0", json:"admin"`
}

func (s *Server) FetchAuth(authD *auth.AuthDetails) (*Auth, error) {
	au := &Auth{}
	err := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (s *Server) DeleteAuth(authD *auth.AuthDetails) error {
	au := &Auth{}
	db := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Server) CreateAuth(userId uint64) (*Auth, error) {
	au := &Auth{}
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	au.AuthUUID = uuid.String()
	au.UserID = userId
	err = s.DB.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (s *Server) IsAdmin(authD *auth.AuthDetails) bool {
	fmt.Println("IsAdmin executing")
	if authD.Admin == 1 {
		return true
	}

	return false
}

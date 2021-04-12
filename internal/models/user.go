package models

import (
	"errors"

	"github.com/HellstromIT/authservice/internal/crypt"
	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `json:"password,omitempty"`
	Admin    int64  `gorm:"default:0" json:"admin"`
}

func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("Email required")
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}

func (s *Server) CreateUser(user *User) (*User, error) {
	if emailErr := s.ValidateEmail(user.Email); emailErr != nil {
		return nil, emailErr
	}

	// Hash password
	hashedPwd, err := crypt.HashAndSalt([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPwd)

	err = s.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	if err := s.DB.Debug().Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

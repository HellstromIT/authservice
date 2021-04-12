package service

import (
	"github.com/HellstromIT/authservice/pkg/auth"
)

type signInInterface interface {
	SignIn(auth.AuthDetails, *string) (string, error)
}

type signInStruct struct{}

var (
	Authorize signInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD auth.AuthDetails, jwt_token *string) (string, error) {
	token, err := auth.CreateToken(authD, jwt_token)
	if err != nil {
		return "", err
	}
	return token, nil
}

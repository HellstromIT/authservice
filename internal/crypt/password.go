package crypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func ComparePasswords(hashedPwd []byte, pwd []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)
	if err != nil {
		log.Println("Failed login attempt")
		return false, err
	}

	return true, nil
}

package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
	Admin    int64
	Expires  int64
}

func CreateToken(authD AuthDetails, jwtSecret *string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUuid
	claims["user_id"] = authD.UserId
	claims["admin"] = authD.Admin
	claims["exp"] = authD.Expires
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(*jwtSecret))
}

func TokenValid(r *http.Request, jwtSecret *string) error {
	token, err := VerifyToken(r, jwtSecret)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken
func VerifyToken(r *http.Request, jwtSecret *string) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method error: %v", token.Header["alg"])
		}
		return []byte(*jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request, jwtSecret *string) (*AuthDetails, error) {
	token, err := VerifyToken(r, jwtSecret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		admin, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["admin"]), 10, 64)
		if err != nil {
			return nil, err
		}
		expires, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AuthDetails{
			AuthUuid: authUuid,
			UserId:   userId,
			Admin:    admin,
			Expires:  expires,
		}, nil
	}
	return nil, err
}

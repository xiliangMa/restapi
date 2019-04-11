package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret []byte = []byte("Hello World！This is jwt test demo!")
)

func GenToken(name, pwd string) (string, int) {
	claims := &jwt.StandardClaims{
		Id:        name,
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "Bitch",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return err.Error(), SiginErr
	}
	return ss, Success
}

func CheckToken(token string) (string, int) {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err.Error(), AuthorizeErr
	}
	return "", Success
}

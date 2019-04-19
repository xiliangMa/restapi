package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)

var (
	secret []byte = []byte("Hello World！This is jwt test demo!")
)

func GenToken(name, pwd string) (string, int) {
	claims := &jwt.StandardClaims{
		Id:        name,
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1800000*6),
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

func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}

func GetNameFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
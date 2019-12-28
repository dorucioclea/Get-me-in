package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenString string)  bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("this is the sample key"), nil
	})

	if token.Valid && err == nil {
		return true
	}
	return false
}
package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claim TokenClaims) string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":claim.Issuer,
		"sub":claim.Subject,
		"aud":claim.Audience,
		"exp":claim.Expiration,
		"nbf":claim.NotBefore,
		"iat":claim.IssuedAt,
		"jti":claim.Id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("this is the sample key"))

	if err != nil {
		return err.Error()
	}
	//debug
	//fmt.Println(tokenString, err)
	return tokenString
}

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
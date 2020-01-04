package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claim *TokenClaims) string{

	claims := &jwt.StandardClaims{
		Audience:  claim.Audience,
		ExpiresAt: claim.Expiration,
		Id:        claim.Id,
		IssuedAt:  claim.IssuedAt,
		Issuer:    claim.Issuer,
		NotBefore: claim.NotBefore,
		Subject:   claim.Subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("this is the sample key"))

	if err != nil {
		return err.Error()
	}

	return tokenString
}

func VerifyToken(tokenString string)  bool {

	// Initialize a new instance of `Claims`
	claims := &jwt.StandardClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
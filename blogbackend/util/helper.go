package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretkey = []byte("Secretyoushouldhide")

func GenerateJWT(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	return claims.SignedString(secretkey)
}

func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

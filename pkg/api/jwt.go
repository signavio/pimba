package api

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func NewToken(bucket, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"bucket":    bucket,
		"timestamp": time.Now().Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method!")
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

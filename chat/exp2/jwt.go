package main

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func ParseToken(tokenStr string) jwt.MapClaims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}
	}

	return claims
}

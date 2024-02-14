package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3NjAwOTc2MDAsInVzZXJuYW1lIjoidW5pc2JhZHJpIn0.Y5bZPWK4rYZpEqvjIvC-3IL6WHwkLL8WTMKfWxYrxGg"
	secretKey := []byte("secret-key")

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		fmt.Println("invalid token")
	}

	log.Default().Println(claims)
}

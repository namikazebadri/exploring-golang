package main

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  "unisbadri",
		"expiredAt": time.Date(2025, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	var secretKey = []byte("secret-key")

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		panic(err)
	}

	log.Default().Println(tokenString)
}

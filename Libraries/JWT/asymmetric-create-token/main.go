package main

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"log"
	"time"
)

// to create private key: openssl genrsa -out id_rsa 4096
// to create public from private key: openssl rsa -in id_rsa -pubout -out id_rsa.pub
func main() {
	var signKey *rsa.PrivateKey

	signBytes, err := ioutil.ReadFile("id_rsa")
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = jwt.MapClaims{
		"user":       "namikazebadri",
		"expired_at": time.Now().Add(time.Minute * 1000),
	}

	token, err := t.SignedString(signKey)

	if err != nil {
		panic(err)
	}

	log.Default().Println(token)
}

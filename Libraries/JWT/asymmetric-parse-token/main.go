package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"log"
)

// to create private key: openssl genrsa -out id_rsa 4096
// to create public from private key: openssl rsa -in id_rsa -pubout -out id_rsa.pub
func main() {
	tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX2F0IjoiMjAyNC0wMi0xNlQyMzoyNjo1My43ODI2NTgrMDc6MDAiLCJ1c2VyIjoibmFtaWthemViYWRyaSJ9.UXfQdT_Hd_4A3rT1UtiBLiR2klXSWnzAhmemhSVjGzH5Ms2WjoZ1pPOmZ-KE321gcTkf-kr_YMUte5o4l7oc_Xvjh74mEvl2KsvJH9q45kQdS_ePNODB_bNa4t_dN2z-gYVH1uTwi5sGvHVO4opQXst6fS9615rLnBseKmlITJa3e7RY-aEbVT5WxhWDt980vZEcVaKCXIHeyXCEgLlNDQ2G4VmKljjvdnyj2HDo5rtG6nQNiYzdzmGP6urWRtupx764KNEOjj30lvW87BuRpLBJwcttOEcxci2eahOuH-w0XSnXgAXYPpfqxIxT8mgLaVDYemnPAhVTWm3REKZarmmZbpaL3Wwb-xsZzIlgmmRGW5CAGpN_Gv5IFJVv5S_yLk_aXjLhkveQ8ia-LHjWMoNkVO3w6DI3lNQA4KQ4naBRnOozhJFtFFbfaazmQXPnIYNJnu52to9gWGXeGJhlcgV0vqeavni856EPkJjCpxo8AFJmm_A08-JoibwGBnWbjTBmnQ_TeYW-BZi8si26juGx4SyZdGg_2BtMnMF9bnBwkfftmzmwjw_0C-m6D0Rln8BzOYCvG3GrQLkH5M2ogoWC2sviw4gD89YZEUwjB2vIYNDxra6w6rvA4vOu8GBuxARb4gJTvegIDjwsrKoCEUMLe8kX52HiN_W5MNh81Jc"

	claims := jwt.MapClaims{}

	var verifyKey *rsa.PublicKey

	verifyBytes, err := ioutil.ReadFile("id_rsa.pub")
	if err != nil {
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return verifyKey, nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		fmt.Println("invalid token")
	}

	log.Default().Println(claims)
}

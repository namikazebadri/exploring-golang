package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Unis Badri"

	if strings.Contains(name, "Badri") {
		fmt.Println("He is Unis Badri")
	} else {
		fmt.Println("He is not Unis Badri")
	}
}

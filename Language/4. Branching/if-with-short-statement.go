package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Unis Badri"

	if contained := strings.Contains(name, "Badri"); contained == true {
		fmt.Println("He is Unis Badri")
	}
}

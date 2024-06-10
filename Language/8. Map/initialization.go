package main

import "fmt"

func main() {
	var age = map[string]int{"Unis Badri": 35}

	age["Namikaze Badri"] = 36

	fmt.Println(age)
}

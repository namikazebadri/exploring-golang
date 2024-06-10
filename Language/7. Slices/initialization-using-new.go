package main

import "fmt"

func main() {
	var names = new([50]string)[0:1]

	names[0] = "Unis Badri"

	fmt.Println(names)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names = []string{2: "Unis Badri", 3: "Namikaze badri"}

	fmt.Println(reflect.ValueOf(names).Kind())
	fmt.Println(names)
}

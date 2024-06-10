package main

import "log"

func main() {
	log.Default().SetPrefix("INFO ")
	log.Default().Println("Hello World")
}

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)

		return
	}

	_, err = conn.Write([]byte("Hello, server!\nHello world!"))

	if err != nil {
		fmt.Println(err)

		return
	}

	errCloseConnection := conn.Close()

	if errCloseConnection != nil {
		return
	}
}

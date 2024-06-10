package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)

		return
	}

	for {
		conn, err := socket.Accept()

		if err != nil {
			fmt.Println(err)

			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Received: %s", buf)
}

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// прослушиваем порт 8080
	listen, err := net.Listen("tcp", "127.0.0.1:808")
	defer listen.Close()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Server is listening...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go server(conn)
	}
}

func server(conn net.Conn) {
	conn.Write([]byte("hello"))
	conn.Close()
}

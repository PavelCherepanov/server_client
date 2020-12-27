package main

import (
	"fmt"
	"net"
	"os"
)

const END_DATA = "[END-DATA]"

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:808")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	conn.Write([]byte("hello" + END_DATA))

	var buffer = make([]byte, 256)
	var message string

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			break
		}
		message += string(buffer[:length])
	}
	fmt.Println(message)

}

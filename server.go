package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const END_DATA = "[END-DATA]"

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
			continue
		}
		go server(conn)
	}
}

func server(conn net.Conn) {
	defer conn.Close()
	var buffer = make([]byte, 256)
	var message string
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			return
		}
		message += string(buffer[:length])
		// оканчивается строчка на END_DATA
		if strings.HasSuffix(message, END_DATA) {
			//  возвращает все кроме конца END_DATA
			message = strings.TrimSuffix(message, END_DATA)
			break
		}
	}

	conn.Write([]byte(strings.ToUpper(message)))

}

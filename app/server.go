package main

import (
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs!!!!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close() // Close the listener when the application exits

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	response := "+PONG\r\n"

	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to connection:", err.Error())
		return
	}

	fmt.Println("Response sent to client")
}

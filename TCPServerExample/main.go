package main

import (
	"fmt"
	"io"
	"net"
	"time"
)


func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
	}

	for {
		fmt.Println("waiting for connection")
		connection, err :=listener.Accept()
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
		}
		io.WriteString(connection, "I see you connected.")
		fmt.Fprintf(connection, "I see you connected!")
		connection.Close()
		time.Sleep(5 * time.Second)

	}
}
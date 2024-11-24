package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// Start listening on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
	}

	for {
		// Wait for a connection
		fmt.Println("waiting for connection")
		connection, err := listener.Accept()
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
		}
		// Send a message to the connected client
		io.WriteString(connection, "I see you connected.")
		fmt.Fprintf(connection, "I see you connected!")
		// Close the connection
		connection.Close()
		// Wait for 5 seconds before accepting the next connection
		time.Sleep(5 * time.Second)
	}
}
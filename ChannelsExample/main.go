package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Create a new channel of type string
	messageChannel := make(chan string)

	fmt.Println(runtime.NumGoroutine())
	
	// Start a new goroutine that sends a message into the channel
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work with a sleep
		messageChannel <- "Hello, Channel!"
	}()
		
	fmt.Println(runtime.NumGoroutine())
	// Receive the message from the channel
	message := <-messageChannel

	// Print the received message
	fmt.Println(message)
}
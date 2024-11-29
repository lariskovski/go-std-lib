package main

import (
	"bufio"
	"fmt"
	"strings"

	// "io"
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
	defer listener.Close()
	for {
		// Wait for a connection
		fmt.Println("waiting for connection")
		connection, err := listener.Accept()
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
		}
		go handle(connection)
	}
}

func handle(conn net.Conn){
	err:= conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Connection timeout")
	}
	defer conn.Close()
	method, uri := request(conn)
	if method == "GET" {
		if uri == "/" {
			getResponse(conn)
		}
	}
	badResponse(conn)
}

func request(conn net.Conn) (string, string){
	i := 0
	method := ""
	uri := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line := scanner.Text()
		// fmt.Println(line)
		if i == 0 {
			// request line
			fields := strings.Fields(line)
			method = fields[0]
			uri = fields[1]
			// fmt.Println("METHOD", method)
			// fmt.Println("URI", uri)
			i++
		}
		if line == "" {
			// headers are done
			break
		}
	}
	return method, uri
}

func getResponse(conn net.Conn){
	body := "Hello World"
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func badResponse(conn net.Conn){
	body := "BAD REQUEST"
	fmt.Fprint(conn, "HTTP/1.1 400 BAD REQUEST\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
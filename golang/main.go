package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer) // This is a blocking call
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing the request")
	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\nHello, World!\r\n")) // This is a blocking call
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on: 1729")
	for {
		log.Println("Waiting for connections")
		// Waiting for the connection to happen - this is a blocking call
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection accepted")
		// Running it as goroutines to make the single-threaded server, multi-threaded
		go do(conn)
	}
}

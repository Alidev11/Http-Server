package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Part 1: The server listens on port 4221 using tcp
	tcpServer, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Found error when creating the server")
		os.Exit(1)
	}
	defer tcpServer.Close()
	log.Println("Server listening on port 4221...")

	
	// Part 2: handle client connection to the server
	for{
		// accept the client connection and write a reply to the client
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Fatalln("Can't connect to server!")
		}

		// Part 3: extract url from request
		buffer := make([]byte, 1024)
		
		// Read request and store it
		n, readErr := conn.Read(buffer)
		if readErr != nil {
			log.Fatalln("Can't read the request")
		}
		request := string(buffer[:n])
		
		// Get the URL path
		reqSplitted := strings.Split(request, " ")	
		fmt.Println("url path:", reqSplitted[1])
		
		// Get the parameter
		params := strings.Split(reqSplitted[1], "/")
		fmt.Println(params[2])
		fmt.Println(len(params[2]))

		// return the response
		response := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(params[2]), params[2],
		)

		_, writeErr := conn.Write([]byte(response))
		if writeErr != nil {
			log.Fatalln("Can't write data to the client!")
		}
		conn.Close()
	}	
}

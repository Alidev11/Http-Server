package main

import (
	// "fmt"
	"log"
	"net"
	"os"
)

func main() {
	// -----------------------
	tcpServer, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalln("Found error when creating the server")
		os.Exit(1)
	}
	defer tcpServer.Close()
	log.Println("Server listening on port 4221...")
	


	// ------------------------
	for{
		// conn contains the request
		conn, err := tcpServer.Accept()
		if err != nil {
			log.Fatalln("Can't connect to server!")
			os.Exit(1)
		}

		urlPath := GetUrlPath(conn)

		Respond(urlPath, conn)
		conn.Close()
	}	
}

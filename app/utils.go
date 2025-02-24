package main;


import(
	"strings"
	"fmt"
	"errors"
	"net"
	"log"
	"os"
)

// checks if url is in the server
func isUrlValid(url string) bool{
	switch url{
	case "/":
		return true
	default:
		return false
	}
}

func getParam(urlPath string) (string, error){
	params := strings.Split(urlPath, "/")
	if len(params)>=3{
		fmt.Print(params[2])
		return params[2], nil
	}
	return "", errors.New("no parameter found")
}

func getUrlPath(connection net.Conn) (string){
	// Part 3: extract url from request
	buffer := make([]byte, 1024)
		
	// Read request and store it
	n, readErr := connection.Read(buffer)
	if readErr != nil {
		log.Fatalln("Can't read the request")
		os.Exit(1)
	}

	request := string(buffer[:n])
	
	// Get the URL path
	reqSplitted := strings.Split(request, " ")
	return reqSplitted[1]
}

func respondParameter(parameter string) (string){
	return fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
		len(parameter), parameter,
	)
}

// main function to return a response
func respond(urlPath string, connection net.Conn){
	parameter, _ := getParam(urlPath)
	response := ""
	
	if parameter == ""{
		if !isUrlValid(urlPath){
			response = "HTTP/1.1 404 NOT FOUND"
		}else{
			response = "HTTP/1.1 200 OK"
		}
	}else{
		response = respondParameter(parameter)
	}

	_, writeErr := connection.Write([]byte(response))
	if writeErr != nil {
		log.Fatalln("Can't write data to the client!")
		os.Exit(1)
	}	
}
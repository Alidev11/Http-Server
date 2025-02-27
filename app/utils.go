package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// ************** Request Methods ********************
func ParseRequest(connection net.Conn) map[int]string{
	buffer := make([]byte, 1024)
	var m map[int]string = make(map[int]string)
		
	n, readErr := connection.Read(buffer)
	if readErr != nil {
		log.Fatalln("Can't read the request")
		os.Exit(1)
	}

	request := string(buffer[:n])
	reqSplitted := strings.Split(request, "\r\n")

	for i, value := range reqSplitted{
		m[i] = value
	}
	fmt.Println(m)

	return m
}

func GetUrlPath(parsedReq map[int]string) string{
	requestLine := parsedReq[0]
	urlPath := strings.Split(requestLine, " ")[1]
	return urlPath
}

func GetParam(urlPath string) (string, error){
	params := strings.Split(urlPath, "/")
	if len(params)>=3{
		return params[2], nil
	}
	return "", errors.New("no paramet found")
}

func ReadHeader(parsedReq map[int]string) (string){
	for _, value := range parsedReq{
		userAgentHeader, cond := strings.CutPrefix(value, "User-Agent: ")
		// fmt.Printf(userAgentHeader)
		fmt.Println(cond)
		if cond{
			return fmt.Sprintf(
				"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
				len(userAgentHeader), userAgentHeader,
			)
		}
	}

	return ""
}

// *************** Response Methods ********************

func Respond(urlPath string, parsedReq map[int]string){
	parameter, _ := GetParam(urlPath)
	response := ""
	
	if parameter == ""{
		switch urlPath{
			case "/":
				response = "HTTP/1.1 200 OK\r\n\r\n"
			case "/user-agent":
				response = ReadHeader(parsedReq)
			default:
				fmt.Println("Default response")
				response = "HTTP/1.1 404 NOT FOUND\r\n\r\n"
		}
	}else{
		response = fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(parameter), parameter,
		)
	}

	_, writeErr := Conn.Write([]byte(response))
	if writeErr != nil {
		log.Fatalln("Can't write data to the client!")
		os.Exit(1)
	}	
}
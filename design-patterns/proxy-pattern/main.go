package main

import "fmt"

func main() {
	server := newNginxServer()

	code, response := server.handleRequest(GetStatus, "GET")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", GetStatus, code, response)

	code, response = server.handleRequest(GetStatus, "GET")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", GetStatus, code, response)

	code, response = server.handleRequest(GetStatus, "GET")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", GetStatus, code, response)

	code, response = server.handleRequest(CreateUser, "POST")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", CreateUser, code, response)

	code, response = server.handleRequest(GetStatus, "GET")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", GetStatus, code, response)

	code, response = server.handleRequest(CreateUser, "POST")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", CreateUser, code, response)

	code, response = server.handleRequest(CreateUser, "POST")
	fmt.Printf("Url: %s\tCode: %d\tResponse: %s\n", CreateUser, code, response)
}

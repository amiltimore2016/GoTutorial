package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting the web server...")
	// Do important thing
	fmt.Println("Server started")
	return port, errors.New("Somthing went wrong")
}

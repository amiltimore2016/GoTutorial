package main

import (
	"fmt"
)

//Channels are a mechanism for communication

func WriteMessageToChannel(message chan string) {
	message <- "hello world"
}

func main() {
	fmt.Println("Channel demo")

	message := make(chan string)

	go WriteMessageToChannel(message)

	fmt.Println("Greeting from the Message Channel: ", <-message)

	close(message)
}

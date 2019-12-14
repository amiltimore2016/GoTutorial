package main

import "fmt"

func main() {
	messageQueue := make(chan string, 4)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"
	messageQueue <- "four"

	close(messageQueue)

	for m := range messageQueue {
		fmt.Println(m)
	}
}

package main

import "fmt"

func main() {
	messageQueue := make(chan string, 6)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"
	messageQueue <- "four"
	messageQueue <- "five"
	messageQueue <- "six"

	close(messageQueue)

	for m := range messageQueue {
		fmt.Println(m)
	}
}

package main

import (
	"fmt"
)

var done chan bool = make(chan bool)

func PrintGreeting(author string) {
	for i := 0; i < 9; i++ {
		fmt.Println("Hello gopher", i, author)
	}
	if author == "goroutine" {
		done <- true
	}

}

// Buffering capacity == 1 a synchronous, gt 1 then its a message queue

func main() {
	go PrintGreeting("goroutine")
	PrintGreeting("main function")

	<-done
}

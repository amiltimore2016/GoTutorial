package main

import (
	"flag"
	"fmt"
)

func main() {
	ArrayValues := [...]int{1, 2, 3}
	var gopherName string
	fmt.Println(ArrayValues[2])
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")
}

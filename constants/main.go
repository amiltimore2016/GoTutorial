package main

import (
	"fmt"
)


// Constant expressions will be evaluate at compile time, not at run time
const (

	first = iota + 6
	second
	third = 2 << iota
)

//When you declare a new constant block it resets the iota count
const (
	third = iota
)

func  main()  {
	fmt.Println(first, second, third)
}
package main

import (
	"fmt"
)

// var SliceName []type = make([]type, # of elements)

func main() {

	var mySlice []int
	fmt.Printf("Contents of mySlice: %v\n", mySlice)
	mySlice = populateIntegerSlice(mySlice, 10)
	fmt.Println("The length of mySlice is: ", len(mySlice))
	fmt.Println("lenght of my slice :", len(mySlice))
	fmt.Println(mySlice)
}

func populateIntegerSlice(input []int, lengthSlice int) []int {

	for index := 0; index < lengthSlice; index++ {
		input = append(input, index+1)
		fmt.Println(input)
	}

	return input
	// We set values in a slice, just like we do with arrays using the square brackets
	// notation [] with the element index enclosed within the square brackets.
}

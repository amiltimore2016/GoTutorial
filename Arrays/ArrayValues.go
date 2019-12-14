// Golang program to illustrate the
// concept of passing a pointer to an
// array as an argument to the function
// using a slice
package main

import "fmt"

// Main Function
func main() {

	// Taking an slice
	s := [5]int{78, 89, 45, 56, 14}
	myMatrix := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	// passing slice to the
	// function updateslice
	updateslice(s[:])

	// displaying the result
	fmt.Println(s)
	printThreeByFourMatrix(myMatrix)
}

func printThreeByFourMatrix(inputMatrix [3][4]int) {
	rowLength := len(inputMatrix)
	colLength := len(inputMatrix[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			fmt.Printf("%5d ", inputMatrix[i][j])
		}
		fmt.Println()
	}
}

// taking a function
func updateslice(funarr []int) {

	// updating the value
	// at specified index
	funarr[4] = 750
}

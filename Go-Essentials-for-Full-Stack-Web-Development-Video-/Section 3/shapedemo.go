package main

import (
	"fmt"
	"github.com/amiltimore2016/GoTutorial/simpleshape"
)

func main() {

	r := simpleshape.NewRectangle(9, 6)
	t := simpleshape.NewTriangle(3, 6)

	fmt.Println("Area of myRectangle: ", simpleshape.ShapeArea(r))
	fmt.Println("Area of myTriangle: ", simpleshape.ShapeArea(t))

}

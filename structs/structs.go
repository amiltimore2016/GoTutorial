package main

//We Declare Methods Using a Pointer Receiver

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(b float64, h float64) *Triangle  {
	return &Triangle{base: b, height: h}
}

func (TriangleP *Triangle Area() float64 {
	return (t.base * t.height)
}
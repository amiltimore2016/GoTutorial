package simpleshape

const mypi = 3.14

type Circle struct {
	radius float64
	pi     float64
}

func NewCircle(r float64) *Circle {
	return &Circle{radius: r, pi: mypi}
}

func (c *Circle) Area() float64 {
	return c.pi * (c.radius * c.radius)
}

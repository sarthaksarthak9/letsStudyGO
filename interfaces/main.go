package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64 // only mention common functionalies, otherwise it will not work
}

type circle struct {
	radius float64
}

type rect struct {
	height float64
	width  float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {

	c1 := circle{4.5}
	r1 := rect{4, 5}

	shapes := []shape{&c1, &r1} // create a slice of type shape. Like if I have a lot of struct having common
	// method like area(), now we can run a for loop and access all struct area()
	// While it is recommended to use pointres when we are working with slice.

	// It is not neccessary to use interface as a slice only, we can also use them as other type also.

	for _, a := range shapes {
		// fmt.Println(shape.area())
		fmt.Println(getArea(a))
	}

	// fmt.Println(shapes[0].area())
	// shapes[1].area()
}

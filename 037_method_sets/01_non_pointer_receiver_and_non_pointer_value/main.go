package main

import (
	"fmt"
	"math"
)

// declaring struct with radius of float type
type circle struct {
	radius float64
}

// declaring interface and making func area of shape type interface
type shape interface {
	area() float64
}

// making a method by passing regular type value to receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// passing interface type shape as parameter to func info
func info(s shape) {
	fmt.Println("area", s.area())
}
func main() {

	c := circle{5}
	info(c)
}

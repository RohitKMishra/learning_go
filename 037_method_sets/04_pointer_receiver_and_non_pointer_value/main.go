// this program will not work

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	c := circle{5}
	info(c) // since the value is not of pointer type so it will not be accepted by a pointer receiver
}

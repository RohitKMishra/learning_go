package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (sq square) area() float64 {
	return sq.length * sq.length
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {

	cir := circle{
		radius: 12.345,
	}
	squ := square{
		length: 13.344,
	}

	info(cir)
	info(squ)
}

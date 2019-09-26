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

// creating a pointer receiver which will receive pointer value only
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}
func main() {
	c := circle{5}
	fmt.Printf("%T\n", &c)
	info(&c) // passing pointer value to pointer method

}

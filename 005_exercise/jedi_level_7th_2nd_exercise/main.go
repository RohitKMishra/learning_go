package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "James"
	// or wc also do it like p.first to change to value as address has to passed to func changeMe
	// p.first = "Jimmy"
}

func main() {

	p1 := person{
		"Robin",
		"Hood",
		30,
	}

	fmt.Println(p1)
	changeMe(&p1) // dereferencing and changing the value stored at that address
	fmt.Println(p1)
}

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// method of type person
func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.last)
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   28,
	}
	p1.speak()
}

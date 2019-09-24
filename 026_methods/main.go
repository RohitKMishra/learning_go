package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// A receiver in go is used to attach a varialbe to any defined function
// receiver attached to func speak
// It enables any value of secretAgent agent type to access func speak
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}
func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}

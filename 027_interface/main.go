// An interface is a collection of method signatures that an Object can implement.
// Hence interface defines the behavior of the object.
// For example, a Dog can walk and bark. If an interface declares method signatures of walk and bark while Dog implements walk and bark methods, then Dog is said to implement that interface.

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

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "  -- the secret agent speak")
}
func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "  --the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrr", h.(secretAgent).first)

	}
	fmt.Println("I was passed into bar", h)

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

	p1 := person{
		first: "Dr",
		last:  " Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}

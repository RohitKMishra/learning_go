package main

import (
	"fmt"
)

func main() {
	s := foo()
	fmt.Println(s)
	b := bar()
	fmt.Printf("%T\n", b) // returns func () as return type is func ()
	a := b()
	fmt.Println(a) // return int as return type is int
}
func foo() string {

	s1 := "Hi"
	return s1
}
func bar() func() int {
	return func() int {
		return 354
	}
}

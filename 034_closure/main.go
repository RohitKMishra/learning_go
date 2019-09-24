package main

import (
	"fmt"
)

func main() {

	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // for a scope of a is till four time
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // for is is only  for two times
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

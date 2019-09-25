//closure example

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	// caling func again to check the scope of variable
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

// foo is a func that return a func having int as return type
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

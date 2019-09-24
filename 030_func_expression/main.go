// Value coming out of a function when assigned to a variable is known as " func expression"

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Playground")

	f := func() {
		fmt.Println("my first func expression ")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)

}

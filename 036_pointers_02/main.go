// When to use pointers

package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before", &x) // referencing to get address of the value stored in memory
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x before", &x) // referencing to get address of the value stored in memory
	fmt.Println("x before", x)
}

// passing memory address as parameter
func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y) // dereferencing to get value from memory address
	*y = 43                     // changing value at that address
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
}

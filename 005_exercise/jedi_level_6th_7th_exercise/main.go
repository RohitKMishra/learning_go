// Assign a function to a variable, then call the function

package main

import (
	"fmt"
)

var g func() = func() {
	fmt.Println("g form outside main")
}

func main() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)

	g = f
	fmt.Printf("%T\tthis is from inside g", g)

}

// call back is when a func is passed into a func as an argument

package main

import (
	"fmt"
)

// func that aacept fun and slice of int as parameter and returns int
func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	// passing func and slice of int as argument to foo and storing it in a variable because it returns int
	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

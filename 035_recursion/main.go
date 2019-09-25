package main

import (
	"fmt"
)

func main() {

	i := factorial(5)
	fmt.Println(i)
	j := loopfact(5)
	fmt.Println(j)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
} // recursion
// it will run till n goes to " 0 "

func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

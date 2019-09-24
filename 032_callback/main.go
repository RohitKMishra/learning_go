package main

import (
	"fmt"
)

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...) // variadic parameter
	fmt.Println(s)
}
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

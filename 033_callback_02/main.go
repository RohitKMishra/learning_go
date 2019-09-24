package main

// Call is passing a function in as an argumnent to a function

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("All numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("Even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Odd numbers", s3)
}

// variadic parameter becomes a slice of int when it enters into function
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v) // gives a slice of int
		}
	}
	return f(yi...) // here f will take a variadic parameter and return int after processing
}
func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v) // gives a slice of int
		}
	}
	return f(yi...) // here f will take a variadic parameter and return int after processing
}

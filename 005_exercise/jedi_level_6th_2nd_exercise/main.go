package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := foo(xi...)
	fmt.Println(s)

	val := bar(xi)
	fmt.Println(val)
}

func foo(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

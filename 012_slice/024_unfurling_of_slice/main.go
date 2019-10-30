package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is ", x)
}

// Variadic parameter must be the last peremeter
// i.e all parameter must be defined before variadic parameter
func sum(x ...int) int {
	fmt.Println(x)

	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("The number at position ", i, "is ", v)
	}
	return sum
}

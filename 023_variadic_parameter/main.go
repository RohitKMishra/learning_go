package main

import "fmt"

func main() {
	x1 := add(1, 2, 3, 4, 5, 6) // to save value that will return from add func
	fmt.Println("The total is ", x1)
}

// here ...int is variadic parameter of type int that accept 0 to n number of parameter
func add(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For index pasiton ", i, "we are adding ", v, "  and now it is ", sum)

	}
	return sum
}

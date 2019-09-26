package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 34, 645, 334, 643, 23, 76, 73}
	xs := []string{"Mmmm", "How's", "That", "For ", " Sorting"}

	fmt.Println(xi)
	sort.Ints(xi) // sort is predefined func used to sort variables in ascending order
	fmt.Println(xi)

	fmt.Println("...............")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

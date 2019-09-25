package main

import (
	"fmt"
)

// defer will execute the function once the containing function is executed
func main() {
	defer foo()
	bar()
}
func foo() {
	defer func() {
		fmt.Println("deferred foo")
	}()
	fmt.Println("This will print after bar :)")
}
func bar() {
	fmt.Println("This will print before foo, wo..hoo...")
}

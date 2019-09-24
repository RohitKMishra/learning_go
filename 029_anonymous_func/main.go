package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Hello from anonymous func")
	}() // these parentheses are needed at the closing braces of the fun to run it

	func(x int) {
		fmt.Println("The meaning of life ", x)
	}(30)
}
func foo() {
	fmt.Println(" Hello from foo")
}

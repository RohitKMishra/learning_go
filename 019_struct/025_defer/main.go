// defer keyword is used to run a block of code or a function when the scope in which the function is defined ends

package main

import "fmt"

func main() {
	defer foo()
	bar()
}
func foo() {
	fmt.Println("This will execute in the end")
}
func bar() {
	fmt.Println("This will execute before foo")
}

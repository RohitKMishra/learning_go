package main

import "fmt"

func main() {
	foo()
	bar("Grinch")          // string passed as an argument to call the function
	s1 := woo("Cindy loo") // variable declared to store the return value
	fmt.Println(s1)
	x, y := mouse("James ", "McDonald")
	fmt.Println(x)
	fmt.Println(y)
	// two variable for two values that are being returned from the function
}

func foo() {

	fmt.Println("Hello from foo")
} // string passed as a parameter
func bar(s string) {
	fmt.Println("Hello,", s)
}
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says "Hello"`)
	b := false
	return a, b
}

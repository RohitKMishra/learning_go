package main

import "fmt"

// declaring a variable
var a int

// defining a variable of type hotdog which have int as it's underlaying type
type hotdog int
var b hotdog

func main(){

	a = 20
fmt.Println("this is a", a)
fmt.Printf("%T\n", a)
fmt.Printf("%v\n", a)
	b = 30
fmt.Println("this is b", b)
fmt.Printf("%T\n",b)
fmt.Printf("%v\n", b)

}
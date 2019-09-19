package main

import "fmt"

// CONSTANT are of two type 
//a. constant that are untyped 
//b. constant that are typed
const a = 42
const b = 42.78
const c = "James Bond"

func main(){

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)

}
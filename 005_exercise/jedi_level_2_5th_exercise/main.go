// Using iota , create 4 constants for the last 4 years
//Print the constant values

package main

import "fmt"

const(
	a= 2015 + iota
	b= 2015 + iota
	c= 2015 + iota
	d= 2015 + iota
)
func main(){

	fmt.Println(a,b,c,d)
}
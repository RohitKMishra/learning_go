package main

import "fmt"
 //USING the short declaration operator , ASSING these VALUES to VARIABLES with the 
 // IDENTIFIERS "x" and "y" and "z"
 // a. 42 
 // b. "James Bond"
 // c. true

 func main(){

	x:= 42
	y:= "James Bond"
	z:= true

	// Printing the values in single print statement 
	fmt.Println(x,y,z)

	// Print the values stored in multiple print statement
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
 }
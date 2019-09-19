package main

import "fmt"
// Part 1 
// At package level scope , assign the following values to three variables 
 // a. for x assign 42
 // b. for y assign "James Bond"
 // c. for z assign true

var x = 42
var y = "James Bond"
var z = true

// Part 2
// use fmt.Sprintf to print the values to one single string 
// Assign the returned value of TYPE string
// using the short declaration operator to a variable with IDENTIFIER

 func main(){

	s := fmt.Sprintf("%v\t%v\t%v\t",x,y,z)
	fmt.Println(s)
 }

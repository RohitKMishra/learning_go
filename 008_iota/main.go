package main

import "fmt"

// iota :- It is a pre declared identifier 
// it increases automatically by one in scope of defined constant
// Once it is out of scope then it again starts with 0 in the next constant scope() 
const(
	    _  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
)

func main(){
	    fmt.Printf("%d\n", kb)
		fmt.Printf("%d\t\t\t%b\n", mb, mb)
		fmt.Printf("%d\t\t%b\n", gb, gb)
	

}


	

		

		
	

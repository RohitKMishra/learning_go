package main

import "fmt"

func main(){

	x:=1 
	// Using for loop without condition
	// BREAK gets you out of the loop 
	// CONTINUE will execute the loop again, leaving all the remaining steps in the loop for that specific iteration
	for{
		x++
		if x > 10 {
			break			
		}

		if x%2 == 0{
			continue
		}
		fmt.Println(x)
		
	}
	
}
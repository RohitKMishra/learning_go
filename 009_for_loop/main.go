package main

import "fmt"

func main(){

	 for x:=1; x<10; x++{
		 if x%2==0{
			 fmt.Println(x)
		 }
	 }
// FOR statement with a single condition
// Using for loop as while loop

y := 3
for y< 6{
	fmt.Println(y)
	y++
} 

// Printing all the ASCII value

	d := 33
	for d < 123 {
		fmt.Printf("%v\t%#U\t%#X\n",d,d,d)
		d++
	}
	fmt.Println("Hello, playground")



}
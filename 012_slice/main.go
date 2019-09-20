package main

import "fmt"

func main(){

// slice are composite literal containing type and values
x := []int{2,3,4,5,6,7} 
fmt.Println(x)
fmt.Println(x[0])
fmt.Println(len(x))
 

// range in for loop for slice
// Here i and v are for index and value of the slice
for i,v := range x{
	fmt.Println(i,v)
}

}
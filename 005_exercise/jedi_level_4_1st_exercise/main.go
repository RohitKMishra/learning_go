package main

import "fmt"

func main(){

	x := [5]int{41,12,43,44,45}
	for i, v := range x{
		fmt.Println(i,v)
		
}
fmt.Printf("%T",x)

}
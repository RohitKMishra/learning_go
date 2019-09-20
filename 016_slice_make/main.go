// use of make built in function

package main

import "fmt"

func main(){
// make (slice,length,capacity)
x:= make([]int,10,14) 
fmt.Println(x)
fmt.Println(len(x))
fmt.Println(cap(x))

x[0]= 12
x[9]= 78
fmt.Println(x)

fmt.Println(x)

x = append(x,23,45,56,78)
fmt.Println(x)
fmt.Println(len(x))
fmt.Println(cap(x))

}
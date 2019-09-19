package main

import "fmt"

type goku int
var x goku
var y int

func main(){

fmt.Println(x)
fmt.Printf("%T\n", x)
x = 42
fmt.Println(x)
fmt.Printf("%v\t",x)
y = int(x)
fmt.Println(y)
fmt.Printf("%T",y)
}
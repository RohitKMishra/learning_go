package main

import "fmt"

var a int 

 func main(){

a = 30

fmt.Printf("%d\t\t%b\t\t%#x\t\t\n",a,a,a)

b := (a << 1)  // short hand declaration

fmt.Printf("%d\t%b\t%#x\t",b,b,b)
 }
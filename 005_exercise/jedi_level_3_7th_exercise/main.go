// Create a program that uses a switch statement with no switch expression specified

package main

import "fmt"

func main(){
// If no switch expression is specified then it takes true as defult expression
// When it get's to any true case it will execute that one

switch {

case 2 == 3: fmt.Println("This is not the line")

case 2 == 0: fmt.Println("This is the line")

default: fmt.Println(" This was about to happen")

}

}
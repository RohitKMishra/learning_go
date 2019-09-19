package main

import "fmt"

func main(){
// In switch , it will print oe execute only first correct statement 
// IF a "fallthrough" is mantioned then only it will take the next one
// Even if the mentioned conditon is not true then too it will be executed after "fallthrough"
	switch{

	case false: fmt.Println("this will not print")
	case (2==4): fmt.Println("this will not print2")
	case (3==3): fmt.Println("this will print")
	fallthrough
	case (5==5): fmt.Println("this will print or not!")
	fallthrough
	case (3==6): fmt.Println("this is dumb print")
	fallthrough
	case (10==9): fmt.Println("this print's")
	fallthrough
	case (8==8): fmt.Println("this too print")
	
	}

}
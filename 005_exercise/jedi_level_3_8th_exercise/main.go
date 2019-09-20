// Create a switch a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favsport"

package main

import "fmt"

func main(){

	favsport := ""

	switch favsport {

	case "Volly ball": fmt.Println("This is not my favourite sport")
	case "Badminton": fmt.Println(" May be this one is ")
	case "Cricket": fmt.Println("This has to be the one")
	default: fmt.Print("This is everyone's favourite")


	}
}
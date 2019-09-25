package main

import ...

// switch on  types
// --normally we switch on value of variable
// --go allows yopu to switch on type of variables

type contact struct {
	greeting string
	name string
}

// switch on type works with interfaces
//we'll learn more about interfaces later
func SwitchOnType(x interface{}){
	switch x.(type) { //this is an assert: asserting, :"x is of this type"

case int:
	fmt.Println("int")
case string:
	fmt.Println("string")
case contact:
	fmt.Println("contact")
	default :
	fmt.Println("unknown")


	}
}
func main(){
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact {"Good to see you,","Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
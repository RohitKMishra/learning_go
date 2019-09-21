package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
func main(){

	p1 := person{
		first: "James",
		last: "Bond",
		age: 82,
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 27,
	}
fmt.Println(p1,p2)
fmt.Println(p1.first, p1.last, p1.age,"\n",p2.first , p2.last,p2.age)
}
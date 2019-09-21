package main

import "fmt"

type person struct {
	first string
	last string 
	favourite_ice_cream_flavour []string
}

func main(){

	p1 := person {
		first: "John",
		last : "Shaw",
		favourite_ice_cream_flavour:[]string {"Blue berry","Strawberry", "martini",},
	}

	p2:= person {

		first : "Bennedict",
		last : "Cena",
		favourite_ice_cream_flavour:[]string {"Chocolate","chocolava","capechino",},
		
	}

fmt.Println(p1.first)
 for i1,v1 := range p1.favourite_ice_cream_flavour{

	fmt.Println(i1,v1)
 }


 fmt.Println(p2.first)
 for i2,v2 := range p2.favourite_ice_cream_flavour{

	fmt.Println(i2,v2)
 }
}irst: "John",
		last : "Shaw",
		favourite_ice_cream_flavour:[]string {"Blue berry","Strawberry", "martini",},
	}

	p2:= person {

		first : "Bennedict",
		last : "Cena",
		favourite_ice_cream_flavour:[]string {"Chocolate","chocolava","capechino",},
		
	}

fmt.Println(p1.first)
 for i1,v1 := range p1.favourite_ice_cream_flavour{

	fmt.Println(i1,v1)
 }


 fmt.Println(p2.first)
 for i2,v2 := range p2.favourite_ice_cream_flavour{

	fmt.Println(i2,v2)
 }
}
package main

import "fmt"

func main(){
 s:= struct {
	 first string
	 friends map[string]int
	 favDrinks []string
 }{
	 first : "james",
	 friends : map[string]int{
		 "MoneyPenny": 132,
		 "Dr no" : 777,
		" M"   : 888,
		},
	 favDrinks: []string{
		 "water",
		 "vodka",
		 "Icetea",
	 },

 }

fmt.Println(s.first)
fmt.Println(s.friends)
fmt.Println(s.favDrinks)

for k, v := range s.friends{
	fmt.Println(k,v)
}

for i, val := range  s.favDrinks{
	fmt.Println(i, val)
}
}


package main

import "fmt"

func main(){

	m := map[string][]string{
		"bond_james": []string{"shaken, not stired","Martinis","Women"},
		"moneypenny_miss": []string{"James bond","Literature","Computer science"},
		"no_dr": []string{"Being evil","Ice cream","Sunsets"},
	}
	m["Fleming_ian"] =[]string{"stakes","cigars","espionage"} 
	// adding new value to map 

	delete(m,"no_dr")  // delete an element form map
 for i1, v1 := range m{
	 fmt.Println("This is the record for",i1)
		 for i2, v2 := range v1{
		 fmt.Println("\t",i2,v2)
	 }
 }
//fmt.Println(m)
}
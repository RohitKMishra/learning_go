package main

import "fmt"

func main(){

	s1 := []string{"james", "Bond","Shaken","Not started"}
	s2 := []string{"Miss","Monneypenny","Helloooo","james"}
	fmt.Println(s1)
	fmt.Println(s2)

	ss :=[][]string{s1,s2} // separate slices within one slice
	fmt.Println(ss)
	as := append (s1,s2...)  // single slice
	fmt.Println(as)

	for i, xs := range ss{		// loops through the main slice ss and prints index
		fmt.Println("record: ",i)
		for j, val := range xs{	// loops through the main slice ss and prints index and value
			fmt.Printf("\t index position : %v \t value: \t %v \n", j,val)
		}
	}
}
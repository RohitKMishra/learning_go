package main

import "fmt"

func main(){

	m := map[string]int{
		"james": 42,
		"Marley": 36,
		"Dillon":52,
	}
	fmt.Println(m)
	fmt.Println(m["Marley"])// send key to the element and it will fetch the vlaue of that key element
	fmt.Println(m["Me"])// if element is not present in map then it will provide zero value for that element

//TO check if the value exist in the map or not
	v, ok := m["Me"]
	fmt.Println(v)
	fmt.Println(ok)

	// ok provides a boolean value so it can be used as condition in if

	if v, ok = m["james"]; ok{
		fmt.Println("This is Awesome")

	// Adding an element to the map
	m["Bob"] = 39

	// Delete element from map
	delete (m,"Dillon")
	fmt.Println(m)

	// check before deleting any element from map
	if v, ok = m["Marley"]; ok{
		fmt.Println("value",v)
		delete (m,"Marley")
	}
	fmt.Println(m)
	}

}
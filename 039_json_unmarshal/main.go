package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":" James","Last":"Bond","Age":32},{"First":"Miss","Last":"MoneyPenny","Age":27}]`
	bs := []byte(s) //Converting json data to slice of byte of string
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}
	var people []person
	// unmarshal takes slice of byte and address of variable to store the data that is to be converted from obtained json
	// It returns error so we need a (" Data Structure") variable to store the error if any occur to print it out later
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}

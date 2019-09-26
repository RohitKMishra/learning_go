package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	s1 := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	bs := []byte(s1)
	fmt.Println(s1)
	fmt.Println("----------------")

	people := []person{}
	// var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(people)

	for i, person := range people {
		fmt.Println("People Number", i)
		fmt.Println("")
		fmt.Println(person.First, person.Last, person.Age)

		// an underscore is used in go if we don't want to give a declare a variable at any given point
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}

	}

}

package main

import "fmt"

func main(){

	s := " You are doing Great"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

// Here bs refers to byte of string which has been obtained from the string after converting it into bs
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",s)
// #U is used to get the UNICODE of characters
	for i := 0; i< len(s); i++ {

		fmt.Printf("%#U\n",s[i])
	}

}
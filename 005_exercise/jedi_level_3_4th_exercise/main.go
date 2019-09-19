package main

import "fmt"

func main(){

	bd := 1991
	for { bd++
		if bd<2020{
			break
		}
		fmt.Println(bd)
		
	}
}
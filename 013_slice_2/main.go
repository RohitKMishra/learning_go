// Slicing the slice

package main

import "fmt"
func main(){

		x := []int{4,6,8,10,18}
		fmt.Println(x) // gives complete slice
		fmt.Println(x[2])//gives value at given position of the slice
		fmt.Println(x[1:])// gives value form postion 1 to end of slice
		fmt.Println(x[1:4])// gives value from 1 to just before 4


}
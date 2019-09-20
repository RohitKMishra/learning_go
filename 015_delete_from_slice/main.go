// delete from a slice

package main

import "fmt"

func main(){

	x:= []int{12,54,69,78,6354}
	y:= []int{5,98,68,45}
	// remove 54 and 69 from x after appending y

	x = append(x,y...)
	fmt.Println(x)

	x = append(x[:1],x[3:]...)
	fmt.Println(x)


}
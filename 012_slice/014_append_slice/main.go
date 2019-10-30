// append add the given value to the end of the slice and returns newly formed slice
 package main
 
 import "fmt"

 func main(){

	x:= []int{7,2,3,8,9}
	fmt.Println(x)
	x = append (x,11,75,78,97)
	fmt.Println(x)

	y:= []int{23,674,745,423}
	x = append(x,y...)  // here ...is variadic parameter which accepts any number of parameter
 // we need to add ... after , in append as it accept ... to complete the function parameter
	fmt.Println(x)
}
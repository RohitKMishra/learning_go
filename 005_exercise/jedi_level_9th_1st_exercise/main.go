package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go model("Inside of model func")
	go sub("Inside of sub")
	wg.Wait()
	fmt.Println("Main terminated")
}
func model(s string) {

	fmt.Println(s)
	wg.Done()
}

func sub(s string) {
	fmt.Println(s)
	wg.Done()
}

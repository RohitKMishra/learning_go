package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	gr := 100
	incrementer := 0

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			//fmt.Println(incrementer)
			v := incrementer
			fmt.Println(v)
			v++
			incrementer = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value is", incrementer)
}

// create a func which return a func , assign the returned fun to a variable, call the returned func

package main

import (
	"fmt"
)

func some() func() int {

	return func() int {
		return 42
	}

}

func main() {

	f := some()
	fmt.Println(f())

}

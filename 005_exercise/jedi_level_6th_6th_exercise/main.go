// Build and use an anonymous func

package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("This is an anonymous func")
	}()
}

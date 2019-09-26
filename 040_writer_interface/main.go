package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, Everybody")
	fmt.Fprintln(os.Stdout, "Hello, Everybody")   // println uses Fprint to print out the string passed to is
	io.WriteString(os.Stdout, "Hello, from here") // Fprintln used Writer string from io package to print the statement

}

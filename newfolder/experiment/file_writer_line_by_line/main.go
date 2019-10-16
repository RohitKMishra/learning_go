package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("/home/cst-pc/go/src/workspace/experiment/file_writer_line_by_line/text.csv")
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
	d := []string{"Welcome to the world of Go.", "Go is an awesome language", "Ttry it one and you will fall for it for sure"}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("/home/cst-pc/go/src/workspace/experiment/file_writer_line_by_line/text.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File appended succesfully")

}

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("/home/cst-pc/go/src/workspace/experiment/file_writer2/bytes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d2 := []byte{104, 101, 108, 123, 104, 125, 124, 134, 146, 124, 104, 116}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

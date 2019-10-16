package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02"))
	date := currentTime.Format("2006-01-02")
	fmt.Printf("%T\n", date)
	fmt.Println(date)
	location := "/home/cst-pc/go/src/workspace/work/"
	filedate := (location + date + " info.log")
	file, err := os.OpenFile(filedate, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "exchange", log.LstdFlags)
	defer file.Close()
	//log.SetOutput(file)
	logger.Println("Logging to go file")

}

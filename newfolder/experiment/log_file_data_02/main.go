package main

import (
	"log"
	"os"
)

type resource struct {
	Status   int         `json:"status"`
	Meta     string      `json:"_meta"`
	Resource interface{} `json:"resource"`
}

type Reso struct {
	Output string `json:"output_data"`
}

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)

	e
}

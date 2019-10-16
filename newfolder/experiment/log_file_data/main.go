package main

import (
	"log"
	"net/http"
	"os"
)

var (
	logger *log.Logger
)

func main() {

	logger = log.New(os.Stdout, "web", log.LstdFlags)

	server := &http.Server{
		Addr:    ":8083",
		Handler: routes(),
	}
	log.Fatal(server.ListenAndServe())
}

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/foo", foo)
	return r
}

func foo(w http.ResponseWriter, r *http.Request) {
	logger.Println("request to foo")
	f, err := os.OpenFile("/home/cst-pc/go/src/workspace/experiment/log_file_data/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	logger = log.New(f, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")

	err = f.Close()
	if err != nil {
		log.Println(err)
	}
}

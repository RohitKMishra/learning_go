package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	logger *log.Logger
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logger = log.New(os.Stdout, "web", log.LstdFlags)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}
	server.ListenAndServe()
	// http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("request to foo")
	// })
	// http.ListenAndServe(":8085", nil)
}

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/foo", foo)

	return r
}

func foo(w http.ResponseWriter, r *http.Request) {
	logger.Println("request to foo")
	fmt.Print(logger)
	f, err := os.OpenFile("test.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger = log.New(f, "perfix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more test to append")
}

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type new struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := new{Title: "Awesome news", News: "Today's headline"}
	t, err := template.ParseFiles("basictemplate.html")
	fmt.Println(err)
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Wow, Go is sweet</h1>")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}

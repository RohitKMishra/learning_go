package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func main() {
	tpl, err := template.ParseFiles("here.where")
	if err != nil {
		log.Fatalln(err)
	}
	tpl, err = tpl.ParseFiles("what.why", "now.then")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "now.then", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "what.why", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

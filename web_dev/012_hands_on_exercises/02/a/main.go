package main

import (
	"log"
	"os"
	"text/template"
)

type CaliforniaHotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type hotels []CaliforniaHotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h := hotels{
		CaliforniaHotel{"Paradise", "Outland", "Ohio", 10012, "Southern"},
		CaliforniaHotel{"Blue Moon", "CountrySide", "Phildelphia", 21004, "Central"},
		CaliforniaHotel{"Sunshine", "Middletown", "Texas", 12340, "Northern"}}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}

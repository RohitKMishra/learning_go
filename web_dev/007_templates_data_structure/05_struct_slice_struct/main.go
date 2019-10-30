package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacture string
	Model       string
	Doors       int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no belief",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "MLK",
		Motto: "It's black nigga",
	}

	f := car{
		Manufacture: "Ford",
		Model:       "F250",
		Doors:       2,
	}

	c := car{
		Manufacture: "Toyota",
		Model:       "Corolla",
		Doors:       4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

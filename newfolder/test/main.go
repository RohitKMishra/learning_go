package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Name struct {
	First string `json:"first"`
}

func main() {
	fmt.Println("Inside main")
	http.HandleFunc("/", user)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", "newuser:password1@/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var name Name

	stmt, err := db.Prepare("INSERT INTO NAME(first) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &name)
	//fmt.Println(name)

	_, err = stmt.Exec(name.First)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(name)

	ch := make(chan string)
	go func() { ch <- name.First }()

	out := <-ch
	//fmt.Println(out)

	var tpl *template.Template
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

	err = tpl.Execute(os.Stdout, out)
	if err != nil {
		log.Fatalln(err)
	}
}

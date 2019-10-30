package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	fmt.Println(out)
}

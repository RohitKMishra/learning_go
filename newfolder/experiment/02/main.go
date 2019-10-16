package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Pword string `json:"pword"`
	ID    int    `json:"id"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signin", getUsers).Methods("GET")
	router.HandleFunc("/signin/{id}", getUser).Methods("GET")
	router.HandleFunc("/signup", createUser).Methods("POST")
	router.HandleFunc("/update/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}

var people []person

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Opening database connection")
	db, err = sql.Open("mysql", "newuser:password1@/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var person person

		err := result.Scan(&person.Fname, &person.Lname, &person.Email, &person.Pword, &person.ID)
		if err != nil {
			panic(err.Error())
		}

		people = append(people, person)
	}
	json.NewEncoder(w).Encode(people)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person person
	db, err := sql.Open("mysql", "newuser:password1@/testdb")
	if err != nil {
		panic(err.Error())
	}

	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM person WHERE id =?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&person.Fname, &person.Lname, &person.Email, &person.Pword, &person.ID)
		if err != nil {
			panic(err.Error())
		}
	}
	defer result.Close()
	json.NewEncoder(w).Encode(person)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var keyVal person
	db, err := sql.Open("mysql", "newuser:password1@/testdb")
	if err != nil {
		panic(err.Error())
	}
	// prepare statement to insert data into database
	fmt.Println("About to create a user")

	stmt, err := db.Prepare("INSERT INTO person (fname, lname, email, pword, id) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Reading data from request")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Getting byte string")

	// keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	fmt.Println(keyVal)
	// fname := keyVal.Fname
	// lname := keyVal.Lname
	// email := keyVal.Email
	// pword := keyVal.Pword
	// id := keyVal.ID

	// id, err := strconv.ParseInt(pid, 0, 64)
	// if err != nil {
	// 	panic(err.Error())
	// }

	fmt.Println("Sending query for execution")
	// slicing through the struct and passing values one on one
	_, err = stmt.Exec(keyVal.Fname, keyVal.Lname, keyVal.Email, keyVal.Pword, keyVal.ID)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("New user has been created")

}
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person person
	params := mux.Vars(r)

	db, err := sql.Open("mysql", "user:password@/testdb")
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("UPDATE person SET fname = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &person)
	newName := person.Fname

	_, err = stmt.Exec(newName, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])

	defer db.Close()
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db, err := sql.Open("mysql", "newuser:password1@/testdb")
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
	defer db.Close()

}

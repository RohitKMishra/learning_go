package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

type ExchangeAssets struct {
	Name     string `json:"name"`
	PrecSave int    `json:"precSave"`
	PrecShow int    `json:"precShow"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port, _ := os.LookupEnv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/{name}", assetDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, router))
}

func assetDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var exchangeAssets ExchangeAssets

	uname, _ := os.LookupEnv("UNAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")

	dbConnection := uname + ":" + password + "@/" + dbname

	logDetails("Opening DB Connection")
	// parameter passed in params will be fetched from address bar
	params := mux.Vars(r)
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err.Error())
	}
	// prepare statement
	logDetails("Preapre query to delete user details")
	stmt, err := db.Prepare("DELETE FROM exchangeAssets WHERE name = ?")
	if err != nil {
		logDetails("Error in prepare query")
		panic(err.Error())
	}

	// reading data into body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logDetails("Error in log data into Body")
		panic(err.Error())
	}

	json.Unmarshal(body, &exchangeAssets)

	logDetails("Sending data to be executed in DB")
	_, err = stmt.Exec(params["name"])
	if err != nil {
		logDetails("Error in executing query")
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with Name = %s was deleted", params["name"])
	defer db.Close()
	logDetails("User deleted successfully")
}

func logDetails(s string) {
	address, _ := os.LookupEnv("ADDRESS")
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	location := address
	filedate := (location + date + "info.log")
	file, err := os.OpenFile(filedate, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error)
	}
	logger := log.New(file, "exchange_delete_user", log.LstdFlags)
	defer file.Close()
	logger.Println(s)
}

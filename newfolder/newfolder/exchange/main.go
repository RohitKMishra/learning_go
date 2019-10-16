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
	router.HandleFunc("/", assetAdd).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))

}

func assetAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var exchangeAssets ExchangeAssets

	uname, _ := os.LookupEnv("UNAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")

	dbConnection := uname + ":" + password + "@/" + dbname
	logDetails("Request to assestAdd")
	logDetails("Opening database connection")

	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		logDetails("unable to open database")
		panic(err.Error())
	}

	// Prepare statement
	stmt, err := db.Prepare("INSERT INTO exchangeAssets (name, precSave, precShow) VALUES(?,?,?)")
	if err != nil {
		logDetails("Error in prepare statement")
		panic(err.Error())
	}
	logDetails("Reading request data")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logDetails("Error in loading data in body")
		panic(err.Error())
	}
	json.Unmarshal(body, &exchangeAssets)

	logDetails("Execute statement")

	_, err = stmt.Exec(exchangeAssets.Name, exchangeAssets.PrecSave, exchangeAssets.PrecShow)
	if err != nil {
		logDetails("Error in executing statement")
		panic(err.Error())
	}
	logDetails("New user created")
	defer db.Close()

}

func logDetails(s string) {

	address, _ := os.LookupEnv("ADDRESS")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02"))
	date := currentTime.Format("2006-01-02")
	fmt.Printf("%T\n", date)
	fmt.Println(date)
	location := address
	filedate := (location + date + " info.log")
	file, err := os.OpenFile(filedate, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "exchange", log.LstdFlags)
	defer file.Close()
	//log.SetOutput(file)
	logger.Println(s)
}

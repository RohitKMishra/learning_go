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

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var db *sql.DB
var err error
var logger *log.Logger
var exchangeAssets ExchangeAssets

type ExchangeAssets struct {
	Name     string `json:"name"`
	PrecSave int    `json:"precSave"`
	PrecShow int    `json:"precShow"`
}

func main() {
	currentTime := time.Now()
	fmt.Println("Hello Universe")

	router := mux.NewRouter()

	router.HandleFunc("/", addUser).Methods("POST")

	message, _ := os.LookupEnv("MESSAGE")
	address, _ := os.LookupEnv("ADDRESS")

	fmt.Println(message)
	fmt.Println(address)
	fmt.Printf("%T\n", address)

	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	log.Fatal(http.ListenAndServe(":8083", router))
}
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uname, _ := os.LookupEnv("UNAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")

	dbconnection := uname + ":" + password + "@/" + dbname

	f, err := os.Create("logs.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err = os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println("Request to assestAdd")
	logger.Println("Opening database connection")

	db, err := sql.Open("mysql", dbconnection)
	if err != nil {
		logger.Println("unable to open database")
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO exchangeAssets (name, precSave, precShow) VALUES(?,?,?)")
	if err != nil {
		logger.Println("Error in prepare statement")
	}
	fmt.Println("Reading data from request")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &exchangeAssets)

	_, err = stmt.Exec(exchangeAssets.Name, exchangeAssets.PrecSave, exchangeAssets.PrecShow)
	if err != nil {
		panic(err.Error())
	}

	logger.Println("New user created")
	defer db.Close()
	defer f.Close()
}

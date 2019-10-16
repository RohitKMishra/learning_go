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

type ExchangeMarkets struct {
	MarketName string `json:"marketName"`
	StockName  string `json:"stockName"`
	StockPrec  int    `json:"stockPrec"`
	MoneyName  string `json:"moneyName"`
	MoneyPrec  int    `json:"moneyPrec"`
	FeePrec    int    `json:"feePrec"`
	MinAmount  string `json:"minAmount"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port, _ := os.LookupEnv("PORT")
	router := mux.NewRouter()

	router.HandleFunc("/", marketAdd).Methods("POST")

	logDetails("Serving on port on" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func marketAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uname, _ := os.LookupEnv("UNAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")

	dbConnection := uname + ":" + password + "@/" + dbname

	var exchangeMarket ExchangeMarkets

	logDetails("Opening db connection")
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		logDetails("Error in opening db connection")
		panic(err.Error())
	}

	logDetails("Preparing statement to insert data")
	stmt, err := db.Prepare("INSERT INTO exchangeMarkets (marketName, stockName, stockPrec, moneyName, moneyPrec, feePrec, minAmount) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logDetails("Error in loading to body")
		panic(err.Error())
	}
	json.Unmarshal(body, &exchangeMarket)

	_, err = stmt.Exec(exchangeMarket.MarketName, exchangeMarket.StockName, exchangeMarket.StockPrec, exchangeMarket.MoneyName, exchangeMarket.MoneyPrec, exchangeMarket.FeePrec, exchangeMarket.MinAmount)
	if err != nil {
		logDetails("Error in executing statement")
		panic(err.Error())
	}
	fmt.Println("New market added successfully")
	logDetails("User added successfully")
	defer db.Close()
}

func logDetails(s string) {

	address, _ := os.LookupEnv("ADDRESS")
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	filedate := (address + date + "info.log")
	file, err := os.OpenFile(filedate, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	logger := log.New(file, "marketAdd", log.LstdFlags)
	defer file.Close()
	logger.Println(s)
}

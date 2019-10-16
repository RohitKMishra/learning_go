package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql tutorial")

	db, err := sql.Open("mysql", "root:password@/people")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")

	insert, err := db.Query("INSERT INTO customers (id,name,prod,amount,address) VALUES (11,'haa ho', 'baigan', 20,'idhre')")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted into customers table")
	defer insert.Close()

}

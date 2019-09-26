package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	// If you get a password and then store it into database in encrypted form
	// bs will get the encrypted password after encription
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	// if you want to check if the stored password is matching the password that is provided by user
	// always store password in encrypted form in the database

	loginpword1 := `password123` // backtick to make it a slice of byte

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpword1))
	if err != nil {
		fmt.Println(" YOU CAN'T LOGIN")
		return
	}
	fmt.Println("YOU ARE LOGGED IN")

}

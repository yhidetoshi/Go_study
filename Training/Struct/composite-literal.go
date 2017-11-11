package main

import "fmt"

type Account struct {
	Email    string
	Password string
	Rank     int
}

func main() {
	account := &Account{
		Email:    "sample@gmail.com",
		Password: "password",
		Rank:     1,
	}
	fmt.Printf("%v", account)
}

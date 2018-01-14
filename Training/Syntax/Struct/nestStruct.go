package main

import "fmt"

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed Feed
}

func main() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Feed.Amount)

}

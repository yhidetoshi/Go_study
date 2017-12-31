package main

import (
	"fmt"
)

const (
	km = 1000000
	m  = 1000
	cm = 10
)

func main() {
	var inputNum int
	var tani string

	fmt.Scan(&inputNum, &tani)

	if tani == "km" {
		fmt.Println(inputNum * km)
	}
	if tani == "m" {
		fmt.Println(inputNum * m)
	}
	if tani == "cm" {
		fmt.Println(inputNum * cm)
	}
}

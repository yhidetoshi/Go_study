package main

import (
	"fmt"
)

const (
	N = 7
)

func main() {

	var input string
	var count int

	for i := 0; i < N; i++ {

		fmt.Scan(&input)
		if input == "no" {
			count++
		}
	}
	fmt.Println(count)
}

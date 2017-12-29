package main

import (
	"fmt"
)

const RATE = 3

func calc(inputN int) {
	if inputN >= 1 {
		fmt.Println(inputN * RATE)
	} else {
		fmt.Println(1)
	}
}

func main() {

	var inputN int
	fmt.Scan(&inputN)
	calc(inputN)

}

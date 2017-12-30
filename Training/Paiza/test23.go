package main

import (
	"fmt"
)

func main() {

	var inputString string
	var inputNum int

	fmt.Scan(&inputString)
	fmt.Scan(&inputNum)

	str := inputString
	
	for i := 0; i < inputNum; i++ {
		fmt.Printf("%s", string(str[i]))
	}
	fmt.Println()
}

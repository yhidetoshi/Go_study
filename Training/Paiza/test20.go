package main

import (
	"fmt"
)

func main() {
	var input string
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Scan(&input)
	//fmt.Printf("%#v\n", strings.Split(str, ""))

	for i := 0; i < len(str); i++ {
		if input == string(str[i]) {
			fmt.Println(i + 1)
		}
	}

}

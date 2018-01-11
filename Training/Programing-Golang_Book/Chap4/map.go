package main

import (
	"fmt"
)

func main() {
	/*
		ages := map[string]int{
			"alice": 20,
			"bob":   30,
		}
		fmt.Println(ages["alice"]) // 20
	*/

	ages := make(map[string]int)
	ages["alice"] = 20
	ages["bob"] = 30
	fmt.Println(ages["alice"]) // 20
}

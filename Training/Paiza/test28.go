package main

import (
	"fmt"
)

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	if n+m < 16 {
		fmt.Println("HIT")
	} else {
		fmt.Println("STAND")
	}
}

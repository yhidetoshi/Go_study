package main

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	if m%n == 0 {
		fmt.Println(n / m)
	} else {
		fmt.Println(n/m + 1)
	}
}

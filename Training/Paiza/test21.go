package main

import (
	"fmt"
)

func main() {

	var m, n int
	//var shou, amari int

	fmt.Scan(&m, &n)
	fmt.Printf("%d %d\n", m/n, m%n)
}

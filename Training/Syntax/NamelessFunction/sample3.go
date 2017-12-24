package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

func main() {
	res := plusAlias(10, 5)
	fmt.Println(res)
}

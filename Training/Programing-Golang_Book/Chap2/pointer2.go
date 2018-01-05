package main

import "fmt"

func main() {
	var p = f()
	fmt.Println(p)
}

func f() *int {
	var v = 1
	return &v
}

package main

import "fmt"

const (
	X = iota
	Y
	Z
)

/*
const (
	X = 1 + iota
	Y
	Z
)
*/

func main() {
	fmt.Println(X, Y, Z)
}

// 0 1 2
// 1 2 3 

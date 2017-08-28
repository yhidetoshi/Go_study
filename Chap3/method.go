package main

import "fmt"

type myType int

func (value myType) println() {
	fmt.Println(value)
}

func main() {
	var z myType = 1234

	z.println()
}

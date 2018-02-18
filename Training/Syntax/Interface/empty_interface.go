package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 40
	describe(i)

	i = "hello"
	describe(i)
}

package main

import "fmt"

func main() {
	var a [3]int
	var i, v int

	for i, v = range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v = range a {
		fmt.Printf("%d\n", v)
	}

}

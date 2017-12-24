package main

import "fmt"

func main() {

	f := func(x, y int) int { return x + y }
	fmt.Println(f(2, 3))
	fmt.Printf("%T\n", func(x, y int) int {return x + y})
}

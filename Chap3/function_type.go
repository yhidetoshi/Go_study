package main

import "fmt"

func main() {
	var f func(int, int) int

	f = func(a int, b int) int {
		return a + b
	}

	fmt.Println(f(1, 2))

	f = multiply
	fmt.Println(f(1, 2))

}

func multiply(x int, y int) int {
	return x * y
}

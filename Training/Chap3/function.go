package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func main() {

	answer := plus(1, 3)

	fmt.Println(answer)
}

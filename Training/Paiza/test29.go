package main

import (
	"fmt"
)

const N = 9

func main() {

	var input int
	fmt.Scan(&input)

	for i := 1; i <= N; i++ {
		if i <= N-1 {
			fmt.Printf("%d ", input*i)
		}
		if i == N {
			fmt.Printf("%d\n", input*i)
		}
	}
}

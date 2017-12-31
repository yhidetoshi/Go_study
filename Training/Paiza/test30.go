package main

import (
	"fmt"
)

const N = 10

func main() {

	var m, n int
	sum := 0

	fmt.Scan(&m, &n)
	fmt.Printf("%d ", m)

	sum = sum + m
	for i := 1; i < N; i++ {

		if i < N-1 {
			fmt.Printf("%d ", sum+n)
			sum += n
		} else {
			fmt.Printf("%d\n", sum+n)
		}

	}
}

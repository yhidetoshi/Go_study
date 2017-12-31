package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	//var result float64

	fmt.Scan(&n)

	result := math.Log10(n) + 1
	fmt.Println(int(result))
}

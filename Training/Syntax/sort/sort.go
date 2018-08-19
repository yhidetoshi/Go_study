package main

import (
	"fmt"
	"sort"
)

func main(){
	fest := []float64{0.3, 0.1, 1.0, 0.7}

	for _, j := range fest{
		fmt.Printf("%f\t", j)
	}

	fmt.Print("\n")
	sort.Float64s(fest)

	for _, f := range fest {
		fmt.Printf("%f\t", f)
	}
}

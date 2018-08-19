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

	sort.Sort(sort.Reverse(sort.Float64Slice(fest)))
	for _, j := range fest{
		fmt.Printf("%f\t", j)
	}

}

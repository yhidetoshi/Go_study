package main

import "fmt"

func main() {
	array1 := [5]float32{}

	array2 := [6]int{1, 2, 3, 4}

	array3 := [...]string{"One", "Two", "Three"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}

package main

import "fmt"

func main() {

	var array1 [1]byte
	var array2 [5]*int
	var array3 [8][3]int64
	var array4 [2]struct{ x, y int }

	// output
	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3), len(array3[0]))
	fmt.Println(len(array4))

}

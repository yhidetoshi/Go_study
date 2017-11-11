package main

import "fmt"

func main() {
	// slice
	s1 := []int{1, 2, 3, 4}

	// slice add element
	s2 := append(s1, 5, 6)

	// slice add slice
	s3 := append(s2, s1...)

	//output
	fmt.Println(s3)

	//apend
	var b1 []byte
	b2 := append(b1, "abc"...)
	fmt.Println(b2)
}

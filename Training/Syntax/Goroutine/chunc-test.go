package main

import "fmt"

func main() {
	source := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(source)

	u := 5

	for i := u; len(source) > 0; {
		if len(source) < u {
			i = len(source)
		}
		target := source[:i]
		source = source[i:]
		fmt.Println(target)
	}
	//fmt.Println(target[0])
}

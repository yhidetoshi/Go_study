package main

import "fmt"

func main() {
	//slice
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	count := copy(dst[2:], src)

	fmt.Println(dst)
	fmt.Println("count:", count)
}

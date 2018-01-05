package main

import "fmt"

func main() {
	var q [3]int = [3]int{1, 2, 3}
	var p [3]int = [3]int{1, 2}
	r := [...]int{1, 2, 3, 4, 5}

	var v int

	for _, v = range q {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("---")
	for _, v = range p {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("---")
	for _, v = range r {
		fmt.Printf("%d\n", v)
	}

}

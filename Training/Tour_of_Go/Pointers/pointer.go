package main

import "fmt"

func main() {
	//i, j := 42, 2701
	i := 42

	p := &i
	fmt.Println(*p)

	*p = 77
	fmt.Println(i)

}

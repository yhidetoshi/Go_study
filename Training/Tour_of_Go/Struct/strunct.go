package main

import "fmt"

type MyData struct {
	X int
	Y int
}

func main() {
	fmt.Println(MyData{1, 2})

	v := MyData{7, 7}
	fmt.Println(v)
	v.X = 9
	fmt.Println(v)
	fmt.Println(v.X)
}

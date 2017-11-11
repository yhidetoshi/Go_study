package main

import "fmt"

type Data struct {
	X int
	Y int
}

func main() {
	v := Data{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
}

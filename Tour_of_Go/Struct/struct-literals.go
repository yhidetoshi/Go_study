package main

import "fmt"

type Data struct {
	X, Y int
}

var (
	v1 = Data{1, 2}
	v2 = Data{X: 1}
	v3 = Data{}
	v4 = Data{Y: 9}
	p  = &Data{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3, v4)
}

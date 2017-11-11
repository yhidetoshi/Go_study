package main

import "fmt"

type Data struct {
	X int
	Y int
}

func main() {
	v := Data{1, 2}
	p := &v
	p.X = 99
	fmt.Println(v)
}

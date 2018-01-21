package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	ps := make([]Point, 5)

	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
}

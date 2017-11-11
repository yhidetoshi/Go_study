package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
*/

func main() {
	p := &Vertex{4, 3}
	v0 := Vertex{4, 3}
	v0.Scale(2)
	fmt.Println(p.X)
//	ScaleFunc(&v0, 10)

	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(p.X)
//	ScaleFunc(&v, 10)

//	p := &Vertex{4, 3}
	p.Scale(2)
	fmt.Println(p.X)
//	ScaleFunc(p, 10)

//	fmt.Println(v, p)
	
}

/*
4
4
8
*/

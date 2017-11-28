package main

import "fmt"

func main() {

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	f := []bool{true, false, true, false, true}
	fmt.Println(f)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5,true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

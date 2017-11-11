package main

import "fmt"

var(
	x int
	y int
	p int
	num int
)

func main(){

	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&p)

	if x%y != 0{
		num = int(x/y) + 1
	}else{
		num = int(x/y)
	}

	fmt.Println(num*p)
}


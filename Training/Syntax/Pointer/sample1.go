package main

import "fmt"

func main() {

	var pointer *int
	var n int = 100
	pointer = &n

	fmt.Println("nのアドレス", &n)
	fmt.Println("pointer値", pointer)

	fmt.Println("nの値", n)
	fmt.Println("pointerの中身", *pointer)

}


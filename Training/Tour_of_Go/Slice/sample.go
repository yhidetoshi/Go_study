package main

import "fmt"

func main(){
	var slice []int = make([]int,2,3)
	fmt.Println(slice)
	// Length
	fmt.Println(len(slice)) // 2
	// Cap
	fmt.Println(cap(slice)) // 3
}
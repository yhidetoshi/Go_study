package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 100
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

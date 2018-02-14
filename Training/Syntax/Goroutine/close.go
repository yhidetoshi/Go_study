package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	close(ch)
	//ch <- 1 // deadloack
	i, hoge := <-ch
	fmt.Printf("%v,%v\n", i, hoge)
}

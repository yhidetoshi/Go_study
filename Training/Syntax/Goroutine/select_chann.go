package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 3)
	ch1 <- 1
	ch2 <- 2

	select {
	case <-ch1:
		fmt.Println("ch1から受信")
	case <-ch2:
		fmt.Println("ch2から受信")
	case ch3 <- 3:
		fmt.Println("ch3へ送信")
	default:
		fmt.Println("ここへは到達しない")
	}
}

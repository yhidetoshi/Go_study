package main

import (
	"fmt"
	"time"
)

func recive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		// 受信できるか確認
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is done")
}

func main() {
	ch := make(chan int, 20)

	go recive("1st goroutine", ch)
	go recive("2nd goroutine", ch)
	go recive("3rd goroutine", ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(3 * time.Second)
}

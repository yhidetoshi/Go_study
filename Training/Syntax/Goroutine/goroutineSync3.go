package main

import (
	"log"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan int, 3) // 第二引数の値がチャンネルに流せる上限の値
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go func(s string) {
			log.Println("Hello", s)
			<- ch //チャンネルから値を取り出す
			wg.Done()}("World" + strconv.Itoa(i))
	}
	wg.Wait()
}
package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchURL(wg *sync.WaitGroup, q chan string) {
	defer wg.Done()
	for {
		url, ok := <-q //closeされると ok=falseになる
		if !ok {
			return
		}
		fmt.Println("Download:", url)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	q := make(chan string, 5)

	// workerを３つ作成
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fetchURL(&wg, q)
	}

	// 同時に３つまでしか処理できない
	q <- "http://www.example.com"
	q <- "http://www.example.net"
	q <- "http://www.example.net/foo"
	q <- "http://www.example.net/bar"
	q <- "http://www.example.net/baz"
	close(q)
	wg.Wait() //すべてのgoroutineが終了するまで待つ

}

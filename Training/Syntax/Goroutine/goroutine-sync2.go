package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync WaitGroupを生成
	wg := new(sync.WaitGroup)
	isFin := make(chan bool)
	wg.Add(2)

	go hoge(isFin, wg)
	go fuga(isFin, wg)

	wg.Wait()
	close(isFin)
	fmt.Println("Fin")
}

func hoge(isFin chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine hogeee")
	}
	wg.Done()
	isFin <- true
}

func fuga(isFin chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine fugaaa")
	}
	wg.Done()
	isFin <- true
}

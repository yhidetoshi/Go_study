package main

import (
	"fmt"
	"time"
)

const N = 1000000000

func main() {
	// sync WaitGroupを生成
	start := time.Now()
	hoge()
	fuga()
	end := time.Now()
	fmt.Println("Finish")

	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

var sumHoge int

func hoge() {
	for i := 0; i < N; i++ {
		//fmt.Println("Goroutine hogeee +++")
		sumHoge++
	}
}

var sumFuga int

func fuga() {
	for i := 0; i < N; i++ {
		//fmt.Println("Goroutine fugaaa ---")
		sumFuga++
	}
}

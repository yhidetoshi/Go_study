package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	CHUNKSIZE = 5 // 何個まとめて処理するか
)

func recive(s []int, ch chan int, wg *sync.WaitGroup) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum

	wg.Done()
	wg.Wait()
}

func main() {
	start := time.Now() //  時間計測

	var chunkNum int // threadの数になる
	var chunkMod int // CHUNKSIZEでの余り
	result := 0

	wg := new(sync.WaitGroup)
	ch := make(chan int)

	inputData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1, 1}

	chunkNum = len(inputData) / CHUNKSIZE
	chunkMod = len(inputData) % CHUNKSIZE

	if chunkMod == 0 {
		chunkNum = chunkNum
	} else {
		chunkNum = chunkNum + 1
	}

	wg.Add(chunkNum)
	for i := CHUNKSIZE; len(inputData) > 0; {
		if len(inputData) < CHUNKSIZE {
			i = len(inputData)
		}

		target := inputData[:i]
		inputData = inputData[i:]
		fmt.Println(target)
		go recive(target, ch, wg)
	}

	// 合計を受け取りし合算
	for j := 0; j < chunkNum; j++ {
		result += <-ch
	}

	time.Sleep(2 * time.Second)

	close(ch)
	fmt.Printf("Total = %v\n", result)

	end := time.Now() // 時間計測
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}

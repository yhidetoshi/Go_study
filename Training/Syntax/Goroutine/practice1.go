package main

import (
	"fmt"
	"sync"
	"time"
)

func getValue() int {
	/* 乱数生成
	     rand.Seed(time.Now().UnixNano())
	   	num := rand.Intn(10)
	*/
	num := 2
	return num
}

// <バケット名を指定して、トータルのオブジェクトサイズを返す>
// → バケット毎のサイズを出力はできるが、値をもらって合計するところができない (Thread単位しか...)
func recive(s []int, ch chan int, wg *sync.WaitGroup) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
	/*
		var sum int
		for {
			// i はバケットのイメージ
			i, ok := <-ch
			// バケットのサイズを計算する関数をここで呼び出す。戻り値はバケットサイズ
			// valu := getBucketSize(bucketname)

			// 受信しなくなると終了
			if ok == false {
				break
			}
			fmt.Printf("i = %v\n", i)
			sum += i
			ch <- sum
			time.Sleep(2 * time.Second)
		}
	*/
	wg.Done()
	wg.Wait()
}

func main() {
	start := time.Now() //  時間計測
	var times, mods int
	wg := new(sync.WaitGroup)

	//bucketsize := 10
	inputdata := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1, 1}

	ch := make(chan int)

	// 複数スレッドで処理する
	//for i := 0; i < 5; i++ {
	u := 5
	times = len(inputdata) / u
	mods = len(inputdata) % u
	if mods == 0 {
		times = times
	} else {
		times = times + 1
	}

	wg.Add(times)
	for i := u; len(inputdata) > 0; {
		if len(inputdata) < u {
			i = len(inputdata)
		}

		target := inputdata[:i]
		inputdata = inputdata[i:]
		fmt.Println(target)
		go recive(target, ch, wg)
	}

	result := 0
	for j := 0; j < times; j++ {
		result += <-ch
	}
	//result1, resutl2, result3 := <-ch, <-ch, <-ch
	time.Sleep(2 * time.Second)

	close(ch)

	fmt.Printf("Total = %v\n", result)

	end := time.Now() // 時間計測
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	// bool型でchannelを作成
	ch := make(chan bool)

	// goroutineを生成して、サブスレッドで処理する
	go func() {
		time.Sleep(2 * time.Second)
		// chに対してtrueを投げる(送信)
		ch <- true
	}()

	// chに受信があったらisFinに返事する
	// 受信があるまで、処理をブロックし続ける(これが同期の仕組み)
	isFin := <-ch

	//chをクローズする
	close(ch)

	//受信した値を出力する
	fmt.Println(isFin)
	fmt.Println("Finished!!")
}

/*
Start!
true
Finished!!
*/

package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int){
	for {
		i, ok := <-ch
		fmt.Println(ok)
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is done")
}

func main() {
	ch := make(chan int, 10)
	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 10{
		ch <- i
		i++

	}
	close(ch)
	time.Sleep(3* time.Second)
}

/* 実行結果
true
true
1st goroutine 0
true
1st goroutine 3
true
1st goroutine 4
true
true
3rd goroutine 2
true
3rd goroutine 6
true
1st goroutine 5
true
1st goroutine 8
true
1st goroutine 9
false
1st goroutineis done
2nd goroutine 1
false
2nd goroutineis done
3rd goroutine 7
false
3rd goroutineis done
*/

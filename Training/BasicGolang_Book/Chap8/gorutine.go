package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("main:start")

	fmt.Println("testを通常の関数として起動")
	test()

	fmt.Println("testをgorutineとして起動")
	go test()

	// 3秒待つ
	time.Sleep(3*time.Second)

	fmt.Println("main:end")
}

func test(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// 1sec wait
		time.Sleep(1*time.Second)
	}
}

package main

import (
	"fmt"
	"log"
)

func main(){
	fin := make(chan bool)

	go func(){
		log.Println("worker working..")
		fmt.Println("worker working..")
		fin <- false
	}()
	<- fin
}
/* 実行結果
2019/01/27 08:45:04 worker working..
worker working..
*/

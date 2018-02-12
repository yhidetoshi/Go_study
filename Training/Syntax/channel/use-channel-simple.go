package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	// チャネル
	sleep1Finished := make(chan bool)
	sleep2Finished := make(chan bool)
	sleep3Finished := make(chan bool)

	go func() {
		//takes 1sec
		log.Print("sleep1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished")
		sleep1Finished <- true
	}()

	go func() {
		//takes 2sec
		log.Print("sleep2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished")
		sleep2Finished <- true
	}()

	go func() {
		log.Print("sleep3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished")
		sleep3Finished <- true
	}()

	// 終わるまで待つ
	<-sleep1Finished
	<-sleep2Finished
	<-sleep3Finished
	log.Print("all finished")
}

/*  output
2018/02/12 13:30:44 started
2018/02/12 13:30:44 sleep3 started
2018/02/12 13:30:44 sleep1 started
2018/02/12 13:30:44 sleep2 started
2018/02/12 13:30:45 sleep1 finished
2018/02/12 13:30:46 sleep2 finished
2018/02/12 13:30:47 sleep3 finished
2018/02/12 13:30:47 all finished
*/

package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	// channel
	finished := make(chan bool)

	go func() {
		//takes 1sec
		log.Print("sleep1 started")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished")
		finished <- true
	}()

	go func() {
		//takes 2sec
		log.Print("sleep2 started")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished")
		finished <- true
	}()

	go func() {
		//takes 3sec
		log.Print("sleep3 started")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished")
		finished <- true
	}()

	// 終わるまで待つ
	<-finished
	<-finished
	<-finished

	log.Print("all finished")
}

/* output
2018/02/12 13:38:49 started
2018/02/12 13:38:49 sleep3 started
2018/02/12 13:38:49 sleep1 started
2018/02/12 13:38:49 sleep2 started
2018/02/12 13:38:50 sleep1 finished
2018/02/12 13:38:51 sleep2 finished
2018/02/12 13:38:52 sleep3 finished
2018/02/12 13:38:52 all finished
*/

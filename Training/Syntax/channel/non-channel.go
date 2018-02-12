package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	// takes 1sec
	log.Print("sleep1 started")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 finished")

	// takes 2sec
	log.Print("sleep2 started")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 finished")

	// takes 3sec
	log.Print("sleep3 started")
	time.Sleep(3 * time.Second)
	log.Print("sleep3 finished")

	log.Print("all finished")
}

/* output
2018/02/12 13:20:53 started
2018/02/12 13:20:53 Sleep1 started
2018/02/12 13:20:54 sleep1 finished
2018/02/12 13:20:54 sleep2 started
2018/02/12 13:20:56 sleep2 finished
2018/02/12 13:20:56 sleep3 started
2018/02/12 13:20:59 sleep3 finished
2018/02/12 13:20:59 all finished
*/

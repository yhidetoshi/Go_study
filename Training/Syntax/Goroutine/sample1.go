package main

import "log"

func main (){

	resc := make(chan int)

	go func(resc chan int){
		log.Println("another goroutine")
		resc <- 42
	}(resc)

	log.Println("the main goroutine recived value", <-resc)

}

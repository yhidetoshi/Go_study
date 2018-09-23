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

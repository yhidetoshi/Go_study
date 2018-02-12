package main

import (
	"fmt"
	"time"
)

func main() {
	isFin1 := make(chan bool)
	isFin2 := make(chan bool)
	isFin3 := make(chan bool)

	fmt.Println("start!!")

	go func() {
		process("A")
		isFin1 <- true
	}()

	go func() {
		process("B")
		isFin2 <- true
	}()

	go func() {
		process("C")
		isFin3 <- true
	}()

	<-isFin1
	<-isFin2
	<-isFin3
	fmt.Println("Fin")
}

func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}

/*
➜ go run subthread2.go
start!!
C
B
A
Fin
➜ go run subthread2.go
start!!
B
C
A
Fin
➜ go run subthread2.go
start!!
A
B
C
Fin
*/

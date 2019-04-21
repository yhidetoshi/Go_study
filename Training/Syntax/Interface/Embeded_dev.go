package main

import (
	"fmt"
)

type Speak interface {
	Hello()
	
}

type A struct {
	Speak
	B
}

type B struct {
  word string
}


func (a *A) Hello() {
	fmt.Printf("hello %s\n", a.word)
	a.word = "hoge"
}

func main() {
	a := &A{}
	a.word = "world"
	a.Hello()
	fmt.Println(a.word)
}


// 実行結果
/*
hello world
hoge
*/

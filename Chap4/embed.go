package main

import "fmt"

type embeded struct {
	i int
}

func (x embeded) doSomething() {
	fmt.Println("test.doSomething()")
}

type test struct {
	embeded
}

func main() {
	var x test
	x.doSomething()

	fmt.Println(x.i)
}

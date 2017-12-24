package main

import "fmt"

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}
func main() {
	/*
	f := returnFunc()
	f()
	*/
	returnFunc()()
}

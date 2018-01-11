package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Greet(msg string) {
	fmt.Printf("%s %v", msg, p.name)
}

func main() {
	p := Person{name: "hide"}
	p.Greet("HI")
}

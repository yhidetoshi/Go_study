package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) Greet(msg string) {
	fmt.Printf("%s %s\n", msg, p.Name)
}

func main() {
	p := Person{Name: "hide"}
	p.Greet("hello!!!")
}

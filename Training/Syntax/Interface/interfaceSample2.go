package main

import "fmt"

type Animal interface {
	Eat()
}

type Human struct {
	Vital string
}

func (h *Human) Eat() {
	h.Vital = "full"
}

func serveFood(a Animal) {
	a.Eat()
}

func main() {
	a := &Human{"Empty"}
	serveFood(a)
	fmt.Printf("vital is %s", a.Vital)
}

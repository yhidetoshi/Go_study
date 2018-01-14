package main

import (
	"fmt"
)

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("Wanwan")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("Nya-nya")
}

func MakeAnimalCry(a Animal) {
	fmt.Println("DoCry!")
	a.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)

	MakeAnimalCry(dog)
	MakeAnimalCry(cat)
}

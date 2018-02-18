package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("wanwan")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("nyanya")
}

func MakeAnimalCry(a Animal) {
	fmt.Println("Cry!!!")
	a.Cry()
}

func main() {
	dog := new(Dog)
	MakeAnimalCry(dog)
}

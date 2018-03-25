package main

import (
	"fmt"
)

type Foo struct{}

func (f *Foo) Name() string {
	return "Foo"
}

type Bar struct{}

func (b *Bar) Name() string {
	return "Bar"
}

type Named interface {
	Name() string
}

func CallName(n Named) {
	fmt.Println(n.Name())
}

func main() {
	foo := new(Foo)
	bar := new(Bar)
	CallName(foo)
	CallName(bar)
}

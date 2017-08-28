package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person
	p1.name = "Bob"
	p1.age = 30

	p2 := Person{age: 31, name: "Hoge"}
	p3 := Person{"Jane", 40}

	p4 := &Person{"Mike", 36}

	// Output
	fmt.Println(p1, p2, p3, p4)

}

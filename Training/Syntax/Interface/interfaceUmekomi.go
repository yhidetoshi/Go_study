package main

import "fmt"

type Gender int

const (
	Female = iota
	Male
)

type Person interface {
	Name() string
	Title() string
}

func New(gender Gender, firstName, lastName string) Person {
	p := &person{firstName, lastName}
	if gender == Male {
		return &male{p}
	} else {
		return &female{p}
	}
}

type person struct {
	firstName string
	lastName  string
}
  
func (p *person) Name() string {
	return p.firstName + " " + p.lastName
}

type female struct {
	*person
}

func (f *female) Title() string {
	return "Ms."
}

type male struct {
	*person
}

func(m *male) Title() string {
	return "Mr."
}

func printFullName(p Person) {
	fmt.Println(p.Title(), p.Name())
}

func main() {
	taro := New(Male, "Taro", "Yamada")
	printFullName(taro)

	hanako := New(Female, "Hanako", "Yamada")
	printFullName(hanako)
}

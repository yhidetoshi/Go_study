package main

import "fmt"

type Person struct{
	Name string
}

func (p Person) Greet(msg string){
	fmt.Println(p.Name)
}

func main(){
	p := Person{Name: "Hide"}
	p.Greet("hi")
}

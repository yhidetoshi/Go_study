package Chapter2

import "fmt"

type Person struct {
	name string
}

func (p Person) Greet(){
	fmt.Println(p.name)
}

func main(){

	p := Person{name: "hide"}
	p.Greet()

}
package main

import "fmt"

func main(){

	b := []string{"a", "b", "c"}
	fmt.Println(&b[0])
	fmt.Println(&b[1])
	fmt.Println(&b[2])
	echo(b)

}

func echo(b []string){
	fmt.Println("------")
	fmt.Println(&b[0])
	fmt.Println(&b[1])
	fmt.Println(&b[2])
}
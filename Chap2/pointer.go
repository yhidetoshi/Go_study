package main

import "fmt"

func main(){
	var ptr *int

	var i int = 12345

	ptr = &i

	//outpu 
	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrの値", ptr)

	fmt.Println("iの値", i)
	fmt.Println(*ptr)
}


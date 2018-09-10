package main

import "fmt"

func main(){
	test1 := []string{"a", "b", "c"}
	print(test1)
}


func print(s []string){
	fmt.Println(s)
}
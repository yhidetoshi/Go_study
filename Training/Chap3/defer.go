package main

import "fmt"

func main(){
	
	fmt.Println("start")

	defer delay()

	fmt.Println("finish")
}

func delay(){
	fmt.Println 
}
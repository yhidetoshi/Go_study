package main

import "fmt"

func main(){
	for i := 0; i < 5; i++ {

		switch i{
			case 0:
				fmt.Println("0")
			case 1, 2:
				fmt.Println("1 or 2")
			default:
				fmt.Println("other")
		}

	}
}
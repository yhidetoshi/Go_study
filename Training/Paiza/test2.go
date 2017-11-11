package main

import "fmt"

func main(){
	var in_n int
	
	fmt.Scan(&in_n)

	if in_n >= 1 && in_n <= 100 {
		for i := 1; i <= in_n; i++ {
			fmt.Print("*")
		}
	}
}


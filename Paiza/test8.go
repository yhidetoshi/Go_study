package main

import (
	"fmt"
)

func setMedal(_first, _second, _third string){

	fmt.Printf("Gold %s\n", _first)
	fmt.Printf("Silver %s\n", _second)
	fmt.Printf("Bronze %s\n", _third)
}


func main(){

	var first, second, third string
	fmt.Scan(&first, &second, &third)
	setMedal(first, second, third)
}

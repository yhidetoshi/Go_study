package main

import (
	"fmt"

)

func main(){

	allInstances := [][]string{}
	instance := []string{
		"t2.nano",
		"ap-northeast-1a",
	}

	fmt.Println(instance[0])
	fmt.Println(instance[1])

	fmt.Println("-----")
	allInstances = append(allInstances, instance)
	fmt.Println(allInstances[0][0])
	fmt.Println(allInstances[0][1])
	
}

/*
t2.nano
ap-northeast-1a
-----
t2.nano
ap-northeast-1a
*/

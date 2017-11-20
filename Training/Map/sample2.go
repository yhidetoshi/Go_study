package main

import "fmt"

func main() {
	var map1 map[string]int = map[string]int{}
	map1["one"] = 1

	fmt.Println(map1)
	fmt.Println(map1["one"])

	map2 := map[int]bool{1: true, 2: false}
	fmt.Println(map2)
	fmt.Println(map2[2])

}

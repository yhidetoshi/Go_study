package main

import "fmt"

func main(){
	var map1 map[string]int = make(map[string]int)

	map1["x"] = 10
	map1["y"] = 20
	fmt.Println(map1["x"])
	fmt.Println(map1["y"])

	map2 := map[string]int{"x": 10, "y": 20}
	fmt.Println(map2)
	fmt.Println(map2["x"])
	fmt.Println(map2["y"])
}
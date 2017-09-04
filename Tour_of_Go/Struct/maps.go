package main

import "fmt"

type Mydata struct {
	Lat, Long float64
}

var m map[string]Mydata

func main() {
	m = make(map[string]Mydata)
	m["Bell Labs"] = Mydata{
		11.1, 22.2,
	}
	fmt.Println(m["Bell Labs"])
}

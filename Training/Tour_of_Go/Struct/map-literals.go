package main

import "fmt"

type Mydata struct {
	Lat, Long float64
}

var m = map[string]Mydata{
	"Bell Labs": Mydata{
		11.1, 22.2,
	},
	"Google": Mydata{
		77.7, 99.9,
	},
}

func main() {
	fmt.Println(m)
}

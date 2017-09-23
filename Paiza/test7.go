package main

import (
	"fmt"
)

type weather string

func (w weather)check_weather(rp int) string{

	if (rp >= 0) && (rp <= 30) {
		w = "sunny"
		return string(w)

	} else if (rp >= 31) && (rp <= 70){
		w = "cloudy"
		return string(w)
	}

	w = "rainy"
	return string(w)
}


func main(){
	var rainy_percent int
	var tenki weather

	fmt.Scan(&rainy_percent)

	fmt.Println(tenki.check_weather(rainy_percent))
}

package main

import (
	"fmt"
)

func main() {
	coins := make(map[string]string)

	coins["JAPAN"]  = "JPY"
	coins["USA"]    = "USD"
	coins["EU"]     = "EUR"
	coins["CHAINA"] = "CNY"

	fmt.Println(coins["JAPAN"])
	fmt.Println(coins["USA"])

	for country, coin := range coins {
		fmt.Println(country, coin)
	}

	coin, exist := coins["JAPAN"]
	if exist {
		fmt.Println("Registered", coin)
	}else{
		fmt.Println("Not Registered")
	}
}

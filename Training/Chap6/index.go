package main

import "fmt"

func main() {
	var date [7]string

	date[0] = "日曜"
	date[1] = "月曜"
	date[2] = "火曜"
	date[3] = "水曜"
	date[4] = "木曜"
	date[5] = "金曜"
	date[6] = "土曜"

	for i := 0; i < len(date); i++ {
		fmt.Print(date[i], " ")
	}
	fmt.Println()

	for _, value := range date {
		fmt.Print(value, " ")
	}

	fmt.Println()

}

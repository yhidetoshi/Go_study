package main

import (
	"fmt"
	"strings"
)

func main() {
	var dislikeN, inputRoom string
	var roomN, counter int
	var result []string

	fmt.Scan(&dislikeN, &roomN)

	// 部屋を入力
	for i := 0; i < roomN; i++ {
		fmt.Scan(&inputRoom)
		if (strings.Contains(inputRoom, dislikeN)) == false {
			result = append(result, inputRoom)
			counter++
		}
	}

	// 結果の出力
	for j := 0; j < len(result); j++ {
		fmt.Println(result[j])
	}
	if counter == 0 {
		fmt.Println("none")
	}

}

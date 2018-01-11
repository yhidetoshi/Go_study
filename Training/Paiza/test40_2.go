package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func strStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

func judgeRoom(_roomN int) {

	for i := 0; i < _roomN; i++ {
		inputRoom := strStdin()
		dislikeN := strStdin()

		if (strings.Contains(inputRoom, dislikeN)) == false {
			result := append(result, inputRoom)
		}
	}
	return result
}

func outputResult(_result ...string) {
	for _, output := range _result {
		fmt.Println(output)
	}
	if len(_result) == 0 {
		fmt.Println("none")
	}
}

func main() {
	//var dislikeN, inputRoom string
	var roomN int
	var result []string

	fmt.Scan(&dislikeN, &roomN)

	result = judgeRoom(roomN)
	// 結果を出力
	outputResult(result...)

}

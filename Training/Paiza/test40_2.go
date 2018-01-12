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

func judgeRoom(_dislikeN string, _roomN int) []string {
	var result []string

	for i := 0; i < _roomN; i++ {
		inputRoom := strStdin()

		if (strings.Contains(inputRoom, _dislikeN)) == false {
			result = append(result, inputRoom)
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
	dislikeN := strStdin()
	var roomN int
	fmt.Scan(&roomN)

	result := judgeRoom(dislikeN, roomN)
	outputResult(result...)

}

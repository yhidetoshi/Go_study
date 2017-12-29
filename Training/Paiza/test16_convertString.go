package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {

	var input, strBefore, strAfter string

	input = StrStdin()

	strBefore = input
	strAfter = strings.Replace(strBefore, "False", "True", -1)
	fmt.Printf("%v\n", strAfter)

}

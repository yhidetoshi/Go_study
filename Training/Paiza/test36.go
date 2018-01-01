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
	var strBefore, strAfter string
	strBefore = StrStdin()
	strAfter = strings.Replace(strBefore, "at", "@", -1)
	fmt.Printf("%v\n", strAfter)
}

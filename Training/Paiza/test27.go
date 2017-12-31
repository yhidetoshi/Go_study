package main

import (
	"fmt"
	"strings"
)

func main() {
	var NGWord string
	var InputWord string
	fmt.Scan(&NGWord, &InputWord)

	//fmt.Println(strings.Contains(InputWord, NGWord))

	if strings.Contains(InputWord, NGWord) == true {
		fmt.Println("NG")
	} else {
		fmt.Println(InputWord)
	}

}

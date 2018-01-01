package main

import (
	"fmt"
)

func main() {
	var inputStringFirst, inputStringSecond string

	fmt.Scan(&inputStringFirst, &inputStringSecond)
	fmt.Printf("%s.%s\n", string(inputStringFirst[0]), string(inputStringSecond[0]))

}

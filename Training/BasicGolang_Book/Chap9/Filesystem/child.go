package main

import (
	"fmt"
	"os"
)

func main() {
	current, _ := os.Getwd()
	fmt.Println(current)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create(`./outputDataFile`)
	if err != nil {
		fmt.Println("ERROR")
	}
	defer file.Close()
	output := "hoggeeeeeeeeeee"
	file.Write(([]byte)(output))
}

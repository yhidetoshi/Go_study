package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fp *os.File
	var err error
	var line int

	if len(os.Args) < 1 {
		fp  = os.Stdin
	} else {
		fp, err = os.Open("./credentials.csv")
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// count line nums
		line ++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(line)
}

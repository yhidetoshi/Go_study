package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile(`./inputData.csv`)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Print(string(data))
}

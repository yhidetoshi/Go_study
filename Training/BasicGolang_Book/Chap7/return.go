package main

import (
	"fmt"
	"os"
)

func main(){
	//file open
	file, err := os.Open("test.txt")

	// openに失敗したか判定
	if err != nil {

		//output error info
		fmt.Println(err.Error())

		os.Exit(1)
	}

	//file close
	file.Close()

	fmt.Println("OK")

}
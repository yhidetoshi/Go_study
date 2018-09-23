package main

import (
	"io/ioutil"
	"os"
)

func main(){
	content := []byte("Hello World")
	err := ioutil.WriteFile("./output.txt", content, 0755)
	if err != nil{
		os.Exit(1)
	}
}

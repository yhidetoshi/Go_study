package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	files, _ := ioutil.ReadDir("/Users/hidetoshi/Go_study/Training/Syntax/Interface")
	for _, file := range files{
		fmt.Println(file.Name())
	}
}

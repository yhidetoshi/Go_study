package main

import (
	"fmt"
	"os"
)

const BUFSIZE = 1024

func main(){
	file, err := os.Open(`./sample-file.csv`)
	if err != nil{
		fmt.Println("Fail")
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil{
			//break
		}
		fmt.Print(string(buf[:n]))
	}
}

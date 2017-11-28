package main

import (
	"os"
	"log"
	"fmt"
)

func main() {

/*
	var cwd string
	var err error
	cwd, err = os.Getwd()
*/

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v" ,err)
	}
	fmt.Println(cwd)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

var(
	inputStrRoles = os.Getenv("GO_ENV")
)

func main() {

	fmt.Printf("%T", inputStrRoles)
	var strRoles []string
	strRoles = strings.Split(inputStrRoles, ",")

	//role := strings.Join(strRoles, ",")
	fmt.Println(strRoles) // develop
	fmt.Println(strRoles[0])
	fmt.Println(strRoles[1])

}
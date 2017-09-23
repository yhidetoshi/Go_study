package main

import (
	"fmt"
	"strings"
)

const (
	aster = "*"
)

func main(){
	var num int
	fmt.Scan(&num)

	fmt.Println(strings.Repeat(aster, num))
}

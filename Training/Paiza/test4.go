package main

import (
	"fmt"
	"strings"
)

/*
const (
	NG = "NG"
)
*/

func main(){
	var in_str, ch_str, out_str string
	fmt.Scan(&in_str, &ch_str)

	out_str = ch_str
	if strings.Contains(ch_str, in_str){
		out_str = "NG"
	}
	fmt.Println(out_str)
}

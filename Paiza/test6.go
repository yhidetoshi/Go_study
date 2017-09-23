package main

import (
	"fmt"
	"strings"
)

const (
	NG = "NG"
)

type word string

func (w word) check_contain(pw word) bool {
	if strings.Contains(string(w), string(pw)){
		return true
	}
	return false
}

func (w word)out_string(pw word) string{
	if w.check_contain(pw) {
		return NG
	}
	return string(w)
}


func main(){
	var ng_word, vali_word word
	fmt.Scan(&ng_word, &vali_word)

	fmt.Println(vali_word.out_string(ng_word))
}

package main

import "fmt"

func main() {
	hoge := [...]string{"a", "b", "", "c"}
	//out := [...]string{}
	out := hoge[:0]
	/*
		i := 0
		for _, str := range hoge {
			if str != "" {
				out[i] = str
				i++
			}
		}
		fmt.Println(out)
	*/
	for _, str := range hoge {
		if str != "" {
			out = append(out, str)
		}
	}
	fmt.Println(out)
}

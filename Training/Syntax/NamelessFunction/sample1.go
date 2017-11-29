/*
nameless function
func([引数]) [戻り値の型] { [関数の本体] }
*/

package main

import (
	"fmt"
)

/*
func decorate(src string) string{
	return "<<" + src + ">>"
}
*/

func main() {
//	f := decorate //関数を変数に格納
	f := func(src string) string { return "<<" + src + ">>" }
	ret := f("Golang")

	fmt.Println(ret)
}

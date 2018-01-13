package main

import "fmt"

func main() {

	/* パターン1 */
	//var variable map[ key_type ] value_type = map[ key_type ] valutetype_

	var languages map[string]string = map[string]string{}
	fmt.Println(languages)
	/*
	  map[]
	*/

	/* パターン2 */
	// パターン1を省略して書くとこうなる
	languages := map[string]string{}

	/* パターン3 */
	// variable := make(map[ key_type ] value_type )

	languages := make(map[string]string)
	// メモリ確保を宣言時に行う場合はこれ
	languages := make(map[string]string, 10)
	fmt.Println(languages)
	/*
		map[]
	*/
	/* 値の代入 */
	//languages["日本"] = "Japanise"
	//fmt.Println(languages["日本"])
	/*
	  Japanise
	*/

}

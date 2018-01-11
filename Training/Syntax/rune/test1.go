package main

import "fmt"

/*
func main() {
	nihongo := "日本語"
	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s¥n", nihongo[:2])
}
*/

/*
nihongo = 日本語
nippon = ��
*/

func main() {
	nihongo := "日本語"

	for pos, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, pos)
	}
	/*
	   	nihongoRune := []rune(nihongo)

	     //size := len(nihongoRune)

	   	fmt.Printf("nihongo = %d characters: \n", size)
	     for i := 0; i < size; i++ {
	   		fmt.Printf("%#U\n ", nihongoRune[i])
	   	}
	   	fmt.Printf("\n")
	*/
}

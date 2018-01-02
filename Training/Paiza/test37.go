package main

import (
	"fmt"
)

const (
	N = 6
)

func main() {

	var pullN int                  // おみくじを引く回数
	var n1, n2, n3, n4, n5, n6 int //選んだ番号を入力
	var a1, a2, a3, a4, a5, a6 int
	//var [pullN]int result
	hitCnt := 0

	fmt.Scan(&n1, &n2, &n3, &n4, &n5, &n6)
	fmt.Scan(&pullN) // おみくじを引く回数を入力
	//inputNum := make([]int, 6, 6)
	inputNum := [6]int{n1, n2, n3, n4, n5, n6}
	//hitCnt := make([]int, pullN)
	result := make([]int, pullN)

	for i := 0; i < pullN; i++ {
		fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6)
		hitNum := [6]int{a1, a2, a3, a4, a5, a6}
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				if inputNum[j] == hitNum[k] {
					hitCnt++
				}
			}
		}
		//fmt.Println(hitCnt)
		result[i] = hitCnt
		hitCnt = 0
	}
	fmt.Println(result)
}

package main

import (
	"fmt"
)

const (
	N = 6
)

func main() {

	var pullN int                                                      // おみくじを引く回数
	var pickNum1, pickNum2, pickNum3, pickNum4, pickNum5, pickNum6 int //選んだ番号を入力
	var a1, a2, a3, a4, a5, a6 int                                     //当選番号入力
	hitCnt := 0

	fmt.Scan(&pickNum1, &pickNum2, &pickNum3, &pickNum4, &pickNum5, &pickNum6)
	fmt.Scan(&pullN) // おみくじを引く回数を入力

	inputNum := [N]int{pickNum1, pickNum2, pickNum3, pickNum4, pickNum5, pickNum6}
	result := make([]int, pullN)

	// 当選確認
	for i := 0; i < pullN; i++ {
		fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6) // 当選番号入力
		hitNum := [N]int{a1, a2, a3, a4, a5, a6}

		for _, preinput := range inputNum {
			for _, input := range hitNum {
				if preinput == input {
					hitCnt++
				}
			}
		}
		result[i] = hitCnt
		hitCnt = 0
	}
	// 結果出力
	for l := 0; l < pullN; l++ {
		fmt.Println(result[l])
	}
}

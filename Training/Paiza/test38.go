package main

import (
	"fmt"
)

const (
	SUBJECT       = 5
	PASSLINE      = 160
	PASSLINETOTAL = 350
)

func main() {

	var studentNum, sumScore, mainScoreSci, mainScoreLang, passCounter int
	var group string
	var inputNum1, inputNum2, inputNum3, inputNum4, inputNum5 int

	fmt.Scan(&studentNum) // テストを受ける人数
	//fmt.Scan(&group)

	for i := 0; i < studentNum; i++ {
		fmt.Scan(&group, &inputNum1, &inputNum2, &inputNum3, &inputNum4, &inputNum5)
		score := [SUBJECT]int{inputNum1, inputNum2, inputNum3, inputNum4, inputNum5}

		sumScore = 0
		mainScoreSci = 0
		mainScoreLang = 0

		for j := 0; j < SUBJECT; j++ {
			sumScore = sumScore + score[j] // 5科目の強謙値を計算

			// 理系2科目の合計
			if j == 1 || j == 2 {
				mainScoreSci = mainScoreSci + score[j]
			}
			// 文系2科目の合計
			if j == 3 || j == 4 {
				mainScoreLang = mainScoreLang + score[j]
			}

		}
		// 合格判定
		if group == "s" && (mainScoreSci >= 160 && sumScore >= 350) {
			passCounter++
		}
		if group == "l" && (mainScoreLang >= 160 && sumScore >= 350) {
			passCounter++
		}
	}
	fmt.Println(passCounter)
}

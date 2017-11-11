package main

import (
	"fmt"
)

const (
	Struckout = 3
)

func judge(_count_strike, _count_ball int) {

	if _count_strike == Struckout {
		fmt.Println("out!")
	} else {
		fmt.Println("fourball!")
	}
}

func main() {

	var throw_balls, i int
	var str string
	var count_strike, count_ball int

	fmt.Scan(&throw_balls)
	ball_type := make([]string, throw_balls)

	for i = 0; i < len(ball_type); i++ {

		fmt.Scan(&str)

		if str == "strike" {
			count_strike++
		}
		if str == "ball" {
			count_ball++
		}
		ball_type[i] = str

		if i != throw_balls-1 {
			fmt.Println(ball_type[i] + "!")
		}

	}
	judge(count_strike, count_ball)
}

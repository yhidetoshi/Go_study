package main

import (
	"fmt"
)

const N = 7

func calcAverage(s1, s2, s3, s4, s5, s6, s7 int) int {
	var ave int
	ave = int((s1 + s2 + s3 + s4 + s5 + s6 + s7) / N)

	return ave
}

func judge(ave, base int) {
	if base <= ave {
		fmt.Println("pass")
	} else {
		fmt.Println("failure")
	}

}

func main() {

	var s1, s2, s3, s4, s5, s6, s7, ave, base int

	fmt.Scan(&s1, &s2, &s3, &s4, &s5, &s6, &s7)
	fmt.Scan(&base)

	ave = calcAverage(s1, s2, s3, s4, s5, s6, s7)
	//fmt.Println(ave)
	judge(ave, base)
}

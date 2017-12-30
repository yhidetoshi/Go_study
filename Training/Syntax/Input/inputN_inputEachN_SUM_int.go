/*
入力回数 Nを入力し、N個数の数の合計を計算する
(ex)
7
2 32 9 40 1 8 13
105
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()

	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x := 0
	for i := 0; i < n; i++ {
		x += nextInt()
	}
	fmt.Println(x)
}


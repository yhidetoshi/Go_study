package main

import (
	"fmt"
)

const SUM = 15

func judge(n1, n2, n3, n4 int) {
	fmt.Println(SUM - (n1 + n2 + n3 + n4))
}

func main() {
	var n1, n2, n3, n4 int
	fmt.Scan(&n1, &n2, &n3, &n4)
	judge(n1, n2, n3, n4)
}

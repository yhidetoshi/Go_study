package main

import (
	"fmt"
)

func judge(n1, n2 int) {
	if n1 > n2 {
		fmt.Println(n1)
	} else if n1 < n2 {
		fmt.Println(n2)
	} else {
		fmt.Println("eq")
	}

}

func main() {

	var n1, n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	judge(n1, n2)
}

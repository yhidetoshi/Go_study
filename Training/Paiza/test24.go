package main

import (
	"fmt"
)

func judge(n, m string) {
	if n == m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {

	var n, m string

	fmt.Scan(&n, &m)
	judge(n, m)

}

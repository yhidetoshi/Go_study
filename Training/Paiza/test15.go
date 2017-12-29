package main

import (
	"fmt"
)

func calc(n, d, e int) {
	if n <= d*e {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

}

func main() {

	var n, d, e int
	fmt.Scan(&n, &d, &e)

	calc(n, d, e)
}

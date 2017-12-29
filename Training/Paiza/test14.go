package main

import (
	"fmt"
)

func calcDay(n1, n2, n3, n4, n5 int) {
	fmt.Println(n2 - n1)
	fmt.Println(n3 - n2)
	fmt.Println(n4 - n3)
	fmt.Println(n5 - n4)

}

func main() {

	var n1, n2, n3, n4, n5 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	fmt.Scan(&n5)

	calcDay(n1, n2, n3, n4, n5)
}

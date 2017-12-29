package main

import (
	"fmt"
)

const (
	FREE   = 6000
	NOFREE = 4000
)

func bugetCalc(drinker, noDrinker int) {
	var chargeSum int
	chargeSum = FREE*drinker + NOFREE*noDrinker
	fmt.Println(chargeSum)
}

func main() {
	var drinker, noDrinker int

	fmt.Scan(&drinker, &noDrinker)
	bugetCalc(drinker, noDrinker)
}

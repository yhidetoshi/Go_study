package main

import (
	"fmt"
)

const (
	J = "J"
	Q = "Q"
)

type fortune string

func (f fortune) judgetelling(_firstcard, _secondcard string) {
	if (_firstcard == J) && (_secondcard == J) {
		_secondcard = Q
	}
	fmt.Println(_firstcard, _secondcard)
}

func main() {
	var firstcard, secondcard string
	var judge fortune

	fmt.Scan(&firstcard, &secondcard)
	judge.judgetelling(firstcard, secondcard)
}

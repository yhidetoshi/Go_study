package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "remove newline")
var seq = flag.String("s", " ", "separatror")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *seq))
	if !*n {
		fmt.Println()
	}
}

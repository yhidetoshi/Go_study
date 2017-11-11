package main

import "fmt"

func main() {
	a, b := 5, 5
	call(a, &b)

	fmt.Println("値渡し", a)
	fmt.Println("ポインタ渡し", b)
}

func call(a int, b *int) {
	a = a + 1
	*b = *b + 1
}

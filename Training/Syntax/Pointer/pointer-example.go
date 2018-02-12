package main

import "fmt"

func showValue(array [4]int) {
	for i, _ := range array {
		array[i] = 0
	}
	fmt.Printf("^^ %v\n", array)
}

func showPointer(array *[4]int) {
	for i, _ := range array {
		array[i] = 0
		fmt.Println(&array[i])
	}
	fmt.Printf("@@ %v\n", *array)
}

func main() {
	array := [4]int{1, 2, 3, 4}
	showValue(array)
	fmt.Printf("== %v\n", array)

	showPointer(&array)
	fmt.Printf("-- %v\n", array)
}

/*
^^ [0 0 0 0]
== [1 2 3 4]
0xc4200181e0
0xc4200181e8
0xc4200181f0
0xc4200181f8
@@ [0 0 0 0]
-- [0 0 0 0]
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 10
	fmt.Println("the value=", m["Answer"])

	m["Answer"] = 20
	fmt.Println("the value=", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value=", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("the value=", v, ok)
}

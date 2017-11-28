package main

import "fmt"

type example struct {
	Name string
}

type examples []*example

func doExample(n string) (r *example) {
	r = new(example)
	r.Name = n
	return r
}

func main() {
	var exs examples
	
	exs = append(exs, doExample("A"))
	exs = append(exs, doExample("B"))

	for _, i := range exs {
		fmt.Println(i.Name)
	}
}

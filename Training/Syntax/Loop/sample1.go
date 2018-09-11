package main

import(
	"fmt"
)
func main(){
	sample := []string{"a", "b", "c"}

	for i, v := range sample{
		fmt.Printf("key = %v  value= %v\n", i, v)
	}
}

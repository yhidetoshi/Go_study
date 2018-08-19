package main

import (
  "fmt"
  "strings"
)

func main(){
  array := []string{"C:", "work", "abc.txt"}
  //array2 := string{"C:", "work", "abc.txt"} error
  fmt.Println(array)
  //fmt.Println(array2)

  fmt.Println(strings.Join(array, "/"))
}

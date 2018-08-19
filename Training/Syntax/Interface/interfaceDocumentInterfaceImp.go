package main

import (
  "fmt"
)

func PrintAll(vals []interface{}) {
  for _, val := range vals {
    fmt.Println(val)
  }
}

func main() {
  names := []string{"one", "two", "three"}

  //PrintAll(names)
  // --> cannot use names (type []string) as type []interface {} in argument to PrintAll

  // 明示的に変換が必要( interface型の配列に格納する)
  vals := make([]interface{}, len(names))
  for i, v := range names {
    vals[i] = v
  }
  PrintAll(vals)
}

/*
one
two
three
*/

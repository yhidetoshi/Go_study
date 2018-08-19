package main

import "fmt"

func PrintAnyType(variable interface{}){
  fmt.Println(variable)
}

func main(){
  intValue := int64(10)
  strValue := "interface"
  PrintAnyType(intValue)
  PrintAnyType(strValue)
}

/* Golangのinterface はどんな型でも受け取ることができる
10
interface
*/

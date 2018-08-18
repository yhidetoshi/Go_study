package main

import "fmt"

func printTypeAndValue(i interface{}){
  fmt.Printf("%v (type:%T)\n", i, i)
}

func main(){
  var i interface{}
  printTypeAndValue(i)

  i = 100
  printTypeAndValue(i)

  i = "Hello"
  printTypeAndValue(i)
}

/*
<nil> (type:<nil>)
100 (type:int)
Hello (type:string)
*/

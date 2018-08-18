package main

import "fmt"

type Book struct {
  Title  string
  Author string
}

func (b *Book) String() string{
  return fmt.Sprintf("%s : %s", b.Title, b.Author)
}

func main(){

  var s fmt.Stringer = &Book{Title: "Golang", Author: "Mike"}
  println(s.String())
}


/*
fmt.Stringer インターフェース

type Stringer interface {
	String() string
}
*/


/*
interface{}


Go 言語では、インタフェースが示すメソッドを実装していれば、
そのインタフェースを備えていると判断されます。
 つまり、上記のような空インタフェースは、すべての型が備えているインタフェースということになります
*/

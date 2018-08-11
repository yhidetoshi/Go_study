package main

import (
	"fmt"
)

type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

type Person struct {
	*Name
	Age int
}

func main() {
	n := &Name{FirstName: "hoge", LastName: "fuga"}
	p := &Person{Name: n, Age: 20}

	fmt.Println(p.String())

  var stringer fmt.Stringer = p
  fmt.Println(stringer.String())

}

/*
- https://golang.org/pkg/fmt/#Stringer

type Stringer interface {
    String() string
}

- https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8

このとき，
1. *Name型は，Stringメソッドを持っているため，
2. fmt.Stringerインタフェースを実装していることになる，
3. 実は*Name型をPerson型に埋め込んだことにより，
4. Person型と*Person型もfmt.Stringerインタフェースを実装したことになる
5. つまり，上記の変数pは以下のようにfmt.Stringer型の変数に代入することができる．
*/

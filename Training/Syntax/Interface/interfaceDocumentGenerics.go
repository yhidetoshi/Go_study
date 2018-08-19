package main

import (
	"fmt"
)

// 全ての型を許容するインターフェースを作っておく
type Any interface{}

// ジェネリクス
type GetValuer interface{
  GetValue() Any
}

// Any型で実装
type Value struct {
  a Any
}

func (v *Value) GetValue() Any{
  return v.a
}


func main() {

  // interfaceで受け取る
  var i GetValuer = &Value{10}
  var s GetValuer = &Value{"vvv"}


  var values []GetValuer = []GetValuer{i, s}

  for _, val := range values {
    fmt.Println(val.GetValue())
  }
}

/*
Go の wikiで紹介されている方法です。
interface{} をそのまま使わず、 Any 型を定義して置き換えているだけですが、
これなら後で Any 型の定義をいじれるのでより良さそうです。
これにより、 Generics 的なことができます
*/

/* 出力結果
10
vvv
*/

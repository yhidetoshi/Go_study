package main

import (
  "fmt"
  "reflect"
)

type User struct {
  Id int "ユーザID"
  Name string "名前"
  Age uint "年齢"
}

func main(){
  u := User{Id: 1, Name: "Taro", Age: 20}

  /* 変数tは reflect.Type型 */
  t := reflect.TypeOf(u)

  /* 構造体の全フィールドを処理するループ */
  for i := 0; i < t.NumField(); i ++ {
    /* i番目のフィールドを取得 */
    f := t.Field(i)
    fmt.Println(f.Name, f.Tag)
  }
}

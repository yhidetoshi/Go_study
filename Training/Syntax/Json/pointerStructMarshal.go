package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
)

type V_Info struct {
  Color    string `json:"color"`
  Weapon   string `json:"weapon"`
}


type Person struct {
    Id        int `json:"id"`
    Name      string `json:"name"`
    Birthday  *string `json:"birthday"`
    V_info    V_Info `json:v_info`
}




func main(){

  bytes, err := ioutil.ReadFile("input-data.json")
  if err != nil {
    fmt.Println("err")
  }
  var person []Person
  if err := json.Unmarshal(bytes, &person); err != nil{
    fmt.Println("err")
  }

  for _, p := range person{
      if p.Birthday != nil {
        fmt.Printf("%d : %s (%s)\n", p.Id, p.Name, *p.Birthday)
      } else{
        fmt.Printf("%d : %s\n", p.Id, p.Name)
      }
  }



}

/* OUTPUT
1: yamada red [07-07]
2: tanaka blue [08-20]
3: suzuki green [09-30]
4: miyamoto yellow [01-11]
0: rei  [] <- JSONにbirthdayフィールドがないのでstring型の初期値の空文字が設定される
*/

          // ↓↓↓↓↓↓↓

/* Birthdayを Pointerで定義すると以下の出力になる
1 : yamada (07-07)
2 : tanaka (08-20)
3 : suzuki (09-30)
4 : miyamoto (01-11)
0 : rei
*/

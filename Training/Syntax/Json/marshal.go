package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

// jsonデコード用に構造体を定義
type Person struct {
  Id        int     `json:"id"`
  Name      string  `json:"name"`
  Birthday  string  `json:"birthday"`
}

func main(){
  bytes, err := ioutil.ReadFile("input-data.json")
  if err != nil {
    fmt.Println("err")
  }
  // jsonデコード
  var persons []Person
  if err := json.Unmarshal(bytes, &persons); err != nil {
        fmt.Println("err")
  }

  for _, p := range persons{
    fmt.Printf("%d : %s`\n", p.Id, p.Name)
  }

}

/* OUTPUT
1 : yamada`
2 : tanaka`
3 : suzuki`
4 : miyamoto`
0 : rei`
*/

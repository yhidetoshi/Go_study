package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
)

type Person struct {
    Id        int `json:"id"`
    Name      string `json:"name"`
    Birthday  string `json:"birthday"`
    V_info    struct {
      Color   string `json:"color"`
      Weapon  string `json:"weapon"`
    } `json:"v_info"`
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
      fmt.Printf("%d: %s %s %s\n", p.Id, p.Name, p.V_info.Color, p.V_info.Weapon)
  }
}

/* OUTPUT
1: yamada red Rang
2: tanaka blue Impact
3: suzuki green Blade
4: miyamoto yellow Collider
0: rei
*/

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

func (p *Person) GetInfo() string{
  if p.Birthday != nil {
      return fmt.Sprintf("%d : %s [%s]", p.Id, p.Name, *p.Birthday)
  }
  return fmt.Sprintf("%d : %s", p.Id, p.Name)
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

  for _, p := range person {
    fmt.Println(p.GetInfo())
  }

}


/*
1 : yamada [07-07]
2 : tanaka [08-20]
3 : suzuki [09-30]
4 : miyamoto [01-11]
0 : rei
*/

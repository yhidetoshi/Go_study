package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:user_name`
	Age  uint   `json:"user_age"`
}

func main() {
	u := User{Id: 1, Name: "Taro", Age: 20}
	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
}

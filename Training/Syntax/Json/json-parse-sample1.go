/*
https://qiita.com/nayuneko/items/2ec20ba69804e8bf7ca3
より学習
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Country struct {
	Name  string       `json:"name"`
	Label []Prefecture `json:"label"`
}

//type graphs struct {
type Prefecture struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Location int    `json:"population"`
}

func main() {

	tokyo := Prefecture{Name: "東京都", Label: "東京", Location: 1}
	saitama := Prefecture{Name: "埼玉県", Label: "さいたま市", Location: 2}
	kanagawa := Prefecture{Name: "神奈川県", Label: "横浜市", Location: 3}
	japan := Country{
		Name:  "日本",
		Label: []Prefecture{tokyo, saitama, kanagawa},
	}

	jsonBytes, err := json.Marshal(japan)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	//	fmt.Println(string(jsonBytes))
	out := new(bytes.Buffer)
	// プリフィックスなし、スペース4つでインデント
	json.Indent(out, jsonBytes, "", "    ")
	fmt.Println(out.String())

}

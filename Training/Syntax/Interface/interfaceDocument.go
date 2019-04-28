package main

import (
	"fmt"
)

// Interfaceを宣言
type Accessor interface {
	GetText() string
	SetText(string)
}

/*
Accessorを満たす実装
Interfaceの持つメソッド群を実装していれば
Interfaceを満たすと言える
明示的な宣言は必要なく、実装と完全に分離している
*/

type Document struct {
	text string
}

func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string) {
	d.text = text
}

func main() {
	var doc *Document = &Document{}
	fmt.Println(&Document{})
	fmt.Println(*doc)

	doc.SetText("hogeee")
	fmt.Println(doc.GetText())

  // Accessor は Interfaceを実装しているので、Accessor 型に代入が可能
	var acsr Accessor = &Document{}
	acsr.SetText("fugaaaaa")
	fmt.Println(acsr.GetText())

}

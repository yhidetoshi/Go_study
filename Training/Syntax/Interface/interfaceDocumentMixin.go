package main

import (
	"fmt"
)

// Interfaceを宣言
type Accessor interface {
	GetText() string
	SetText(string)
}


type Document struct {
	text string
}

type Page struct {
  Document
  Page int
}

func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string) {
	d.text = text
}

func main() {
	var doc *Document = &Document{}
	doc.SetText("hogeee")
	fmt.Println(doc.GetText())

  // Accessor は Interfaceを実装しているので、Accessor 型に代入が可能
	var acsr Accessor = &Document{}
	acsr.SetText("fugaaaaa")
	fmt.Println(acsr.GetText())

  var acsr2 Accessor = &Page{}
  acsr2.SetText("mixinnnnnnn")
  fmt.Println(acsr2.GetText())

}

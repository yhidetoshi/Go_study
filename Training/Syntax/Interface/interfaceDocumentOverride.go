package main

import (
	"fmt"
  "strconv"
)

// Interfaceを宣言
type Accessor interface {
	GetText() string
	SetText(string)
}



type Document struct {
	text string
}

type ExtendedPage struct {
  Document
  Page int
}

// Document.GetText()のオーバーライド
func (ep *ExtendedPage) GetText() string{
  return strconv.Itoa(ep.Page) + " : " + ep.Document.GetText()
}


func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string) {
	d.text = text
}

func SetAndGet(acsr Accessor) {
  acsr.SetText("accessor")
  fmt.Println(acsr.GetText())
}


func main() {
  var acsr Accessor = &ExtendedPage{
    Document{},
    2,
  }
  acsr.SetText("Pageee")
  fmt.Println(acsr.GetText())
}


/*
2 : Pageee
*/

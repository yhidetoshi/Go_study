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

func SetAndGet(acsr Accessor) {
  acsr.SetText("accessor")
  fmt.Println(acsr.GetText())
}


func main() {
  SetAndGet(&Page{})
  SetAndGet(&Document{})

}

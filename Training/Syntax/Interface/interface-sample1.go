package main

import "fmt"

type Accessor interface {
	GetText() string
	SetText(string)
}

type Document struct {
	text string
}

func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string){
	d.text = text
}

func main() {
	var doc *Document = &Document{}
	doc.SetText("document")
	fmt.Println(doc.GetText())

	var acsr Accessor = &Document{}
	acsr.SetText("accessor")
	fmt.Println(acsr.GetText())

}

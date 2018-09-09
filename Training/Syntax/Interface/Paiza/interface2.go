package main

import "fmt"

const (
	FIRSTFARE = 100
	SECTIONFARERATE = 10
)

type Calc interface{
	calcFare(int, int)
	setSection() int
	setFirstFare() int
}

type Kipu struct{
	firstFare int
	sectionFare int
}



func main(){
	//var c Calc = &Kipu{}
	var k *Kipu = &Kipu{}
	_sectionFare := k.setSection()
	_firstFare := k.setFirstFare()
	k.calcFare(_firstFare,_sectionFare)
}

func (k *Kipu) setSection() int{
	var _section int
	fmt.Scan(&_section)
	k.sectionFare = _section * SECTIONFARERATE

	return k.sectionFare
}

func (k *Kipu) setFirstFare() int{
	k.firstFare = FIRSTFARE

	return k.firstFare
}

func (k *Kipu) calcFare(f, s int){
	fmt.Println(f + s)
}
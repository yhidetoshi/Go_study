package main

import(
	"fmt"
)

type Calc interface{
	GetSalary(int) int
	GetMultiNum(int) int
	GetBonus(int, int)
}

type Info struct{
	Calc
	saraly int
	multiNum int
	bonus int
}

func main(){
	var c Calc = &Info{}
	var _saraly, _multiNum int

	fmt.Scan(&_saraly, &_multiNum)
	valSaray := c.GetSalary(_saraly)
	valmultiNum := c.GetMultiNum(_multiNum)
	c.GetBonus(valSaray,valmultiNum)
}

func (i *Info) GetSalary(s int) int{
	i.saraly = s
	return i.saraly
}

func (i *Info) GetMultiNum(m int) int{
	i.multiNum = m
	return m
}

func (i *Info) GetBonus(s, m int){
	i.bonus = s * m
	fmt.Println(i.bonus)
}
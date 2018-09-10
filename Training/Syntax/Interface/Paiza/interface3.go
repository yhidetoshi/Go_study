package main

import "fmt"

const (
	S = 'S'
	W = 'W'
	N = 10
	THRESHVALUE = 5
)

type Info interface{
	GetInfo() []string
	Judge([]string)
}

type Clothes struct {
	clothe []string{}
}

func main(){
	var c *Clothes = &Clothes{}

	_result := c.GetInfo()
	c.Judge(_result)
}

func (c *Clothes)GetInfo() []string{
	var _input string

	//fmt.Scan(&_input)
	for i := 0; i < N; i++{
		fmt.Scan(&_input)
		c.clothe[i] = _input
	}
	fmt.Println(len(c.clothe))
	return c.clothe
}

func (c *Clothes) Judge(s []string){
	var counter int
	for i, v := range s{

		if v[i] == W {
			counter ++
		}
	}
	if counter >= THRESHVALUE {
		fmt.Println("OK")
	}
	fmt.Println("NG")
}
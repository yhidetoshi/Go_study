package main

import "fmt"

type Dog struct {}

func (d *Dog) Cry (){
	fmt.Println("Wanwan")
}



type Cat struct {}

func (c *Cat) Cry(){
	fmt.Println("Nya-nay-")
}


func MakeDogCry(d *Dog) {
	fmt.Println("doCry!")
	d.Cry()
}
func (d *Dog)MakeDogCry2() {
	fmt.Println("doCry! 2")
	d.Cry()
}


func MakeCatCry(c *Cat) {
	fmt.Println("doCry")
	c.Cry()
}
func (c *Cat)MakeCatCry2(){
	fmt.Println("doCry2")
	c.Cry()
}


func main(){
	dog := new(Dog)
	var dog2 = Dog{}

	cat := new(Cat)
	var cat2 = Cat{}

	MakeDogCry(dog)
	dog2.MakeDogCry2()

	MakeCatCry(cat)
	cat2.MakeCatCry2()
}

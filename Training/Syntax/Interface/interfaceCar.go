package main

import (
	"fmt"
	"strconv"
)

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走る"
}

func (u *MyCar) stop() {
	fmt.Println("停止する")
	u.speed = 0
}

func main() {

	myCar := &MyCar{name: "マイカー", speed: 0}
	var i Car = myCar
	fmt.Println(i.run(50))
	i.stop()

}

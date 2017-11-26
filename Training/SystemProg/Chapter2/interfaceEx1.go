package Chapter2

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk(){
	fmt.Printf("My name is %s", g.name)
}

func main(){
	var talker Talker
	talker = &Greeter{"hide"}
	talker.Talk()
}
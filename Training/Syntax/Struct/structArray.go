package main

import(
	"fmt"
)

type Animal struct {
	Name string
	Age int
}

// Animalの構造体をもつ配列
type Animals []Animal

func main() {
	// Animal構造体の初期化
	var dog Animal = Animal {
		Name: "Wanwan",
		Age: 5,
	}
	 cat := Animal {
		Name: "nyanya",
		Age: 10,
	}
	
	// Animalの構造体を格納する配列を宣言
	var animals Animals
	
	animals = append(animals, dog)
	animals = append(animals, cat)
	
	fmt.Println(animals)
	
}

/*
[{Wanwan 5} {nyanya 10}]
*/

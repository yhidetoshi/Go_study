package main

import "fmt"

/* interfaceを使う */
type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("にゃーにゃー")
}

/* 鳴かせる部分が冗長 !!
func MakeDogCry(d *Dog) {
	fmt.Println("鳴け！")
	d.Cry()
}

func MakeCatCry(c *Cat) {
	fmt.Println("鳴け！")
	c.Cry()
}
*/

/* 鳴く部分のinterfaceで記述 */
func MakeAnimalCry(a Animal) {
	fmt.Println("CRY!!!")
	a.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	MakeDogCry(dog)
	MakeCatCry(cat)
}

/* README.md
http://cuto.unirita.co.jp/gostudy/post/interface/
ダックタイピング
Animalインタフェースを作成した際、Dog構造体やCat構造体には何も変更を加えていません。 にも関わらず、DogやCatはMakeAnimalCry関数において「Animalインタフェースを実装した」とみなされています。
インタフェースの実装に明示的な宣言が必要なJavaやC#と違い、 Go言語ではインタフェースが必要とするメソッドをすべて実装した時点で、自動的にそのインタフェースを実装したとみなされます。
*/

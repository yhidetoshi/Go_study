# Memo

### 構造体の初期化パターン
- パターン
  - 宣言後にフィールド設定
  - 順番で初期化
  - Key:Valueで初期化
  - newで初期化 u5はポインタ型
  - 初期化関数を利用して初期化



- new()による割り当て
  - (Go風な表現) new(T)が返す値は、新しく割り当てされた型Tのゼロ値のポインタ
  - (一般的な表現) new(T) 型Tの新しいアイテム用にゼロ化した領域を割り当て、そのアドレスである *T型の値を返す
  - new()から返されるメモリはゼロ化されている。→ すぐに使える状態になる

- make()による割り当て
  - makeで割り当てできるのは、 スライス・マップ・チャネル であり、初期化、ゼロ値ではない、T型(*Tでない)の値を返す

```
package main

import "fmt"

//User構造体定義。大文字から始まっているので他パッケージからアクセス可能
type User struct {
	name string
	age  int
}

func main() {

// 宣言後にフィールド設定
	var u User
	u.name = "hide"
	u.age = 20
	fmt.Printf("name:%s", u.name)

// 順番で初期化
	u2 := User{"hide", 20}
	fmt.Printf("name:%s", u2.name)

// Key:Valueで初期化
	u3 := User{name: "hide", age: 30}
	fmt.Println("name:%s", u3.name)
	// Key:Valueで初期化 u4はポインタ型
	u4 := &User{name: "hide", age: 20}
	fmt.Println("name:%s", u4.name)

// newで初期化 u5はポインタ型
	u5 := new(User)
	u5.name = "hide"
	fmt.Println("name:%s", u5.name)

// 初期化関数を利用して初期化
	var user *User = newUser("hide", 20)
	fmt.Printf("name:%s", user.name)
}

// 初期化関数
func newUser(name string, age int) *User{
	u := new(User)
	u.name = name
	u.age  = age
	return u
}
```

### 構造体配列
```
package main
import "fmt"

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
```

### 構造体in構造外(ネスト)
```
type Feed struct {
	Name   string
	Amount uint
}
type Animal struct {
	Name string
	Feed Feed
}
func main() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Feed.Amount)
}
```

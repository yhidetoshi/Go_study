# Memo

# 構造体の初期化パターン
- パターン
  - 宣言後にフィールド設定
  - 順番で初期化
  - Key:Valueで初期化
  - newで初期化 u5はポインタ型
  - 初期化関数を利用して初期化

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

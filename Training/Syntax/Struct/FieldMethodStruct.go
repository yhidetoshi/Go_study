package foo

type T struct {
  Field1 int // 公開フィールド
  field2 int // 非公開フィールド
}

// 公開メソッド
func (t *T) Method1() int{
  return t.Field1
}

// 非公開メソッド
func (t *T) method2() int{
  return t.field2
}

package main

t := &foo.T{}

t.Method1() // OK
t.Field1    // OK

t.method2() // コンパイルNG
t.field2    // コンパイルNG

}

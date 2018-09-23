package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for i :=0; i<3; i++{
		wg.Add(1)
		go func(i int){
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
上記の例では、関数を定義してから、その関数をgoで呼び出しています。
しかし、Go言語では無名関数（クロージャ）が作れるので、
次のように関数の作成とgoroutine化を同時に行うことができます。
この場合はgoの後ろには関数名ではなく「関数呼び出し文」がくるので、
末尾に「()」が必要です。 メソッド呼び出しも使えます。
*/
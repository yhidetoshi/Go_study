#  Gorutine

## チャネル

- channel
  - 1つのgoroutineが他のgorutineへ値を送ることができるようにする仕組み
  - 送受信が完了するまでブロックする
  - チャネル作成
    - `ch := make(chan int)`

- 送信 と 受信 の二つの主要な操作
```
ch <- x   // 送信文
x = <-ch  // 代入文での受信式
<-ch      // 受信文:  結果は破棄
```
- クローズ
  - そのチャネルに値がこれ以上は送信されないことを示すフラグを設定します。その後に送信を試みるとパニックになる。
  - チャネルの要素型のゼロ値を生成する
  - `close(ch)`
  
- バッファ（無し・有り）
```
ch = make(chan int)    // バッファなしチャネル
ch = make(chan int, 0) // バッファなしチャネル
ch = make(chan int, 3) // 容量が3であるバッファありチャネル
```

### WaitGroup
- syncパッケージのWaitGroup

| メソッド | 説明 | 
|:----:|:-----:|
| Add() | WaitGroup のカウンタを上げる|
| Done() | WaitGroup のカウンタを下げる|
| Wait() | WaitGroup のカウンタが0になるまで待つ|

- サンプルコード
```
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
```

```
上記の例では、関数を定義してから、その関数をgoで呼び出している。
しかし、Go言語では無名関数（クロージャ）が作れるので、
次のように関数の作成とgoroutine化を同時に行うことができる。 
この場合はgoの後ろには関数名ではなく「関数呼び出し文」がくるので、
末尾に「()」が必要になる。 メソッド呼び出しも使える。
```

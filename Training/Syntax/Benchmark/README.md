# Benchmark

- Helloworkdでベンチーマクする例
```
package bench

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello world")
	}
}
```

- `$ go test -bench Hello`
  - 入力して実行すると、Helloに一致するベンチマークが実行されます。
- `go test -bench .`



```
type BenchmarkResult struct {
        N         int           // 反復回数
        T         time.Duration // 実行時間の合計
        Bytes     int64         // 1反復あたりのバイト数
        MemAllocs uint64        // メモリ割当量の合計
        MemBytes  uint64        // メモリ割当バイト数の合計
}
```

// http.Get 		HTTPリクエストを発行し、エラーがなければレスポンス構造体で結果を返す
// ioutil.ReadAll 	バッファリングしながらデータ終端までデータを読み込む
// Fprintf(F 接頭辞)     書き込み先を明示的に指定できる

package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)

		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

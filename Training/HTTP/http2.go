package main

import (
	"fmt"
	"net/http"
)

func main(){
	db := database{"shoes":50, "socks":5}
	http.ListenAndServe("localhost:8000", db)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// メソッドのレシーバでdbを受けているので、ServeHTTPを満たしている
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page %s\n", req.URL)
	}
}

/* 実行結果

ブラウザで　
localhost:8000/list
shoes: $50.00
socks: $5.00

http://localhost:8000/price?item=socks
$5.00

http://localhost:8000/price?item=shoes
$50.00

http://localhost:8000/price?item=hat
no such item: "hat"

http://localhost:8000/help
no such page /help
*/


/*
- ハンドラはURLのパスの要素 req.URL.Pathに基づいてどのロジックを実行するかを決める
- ハンドラがパスを認識できなければ、w.WriteHeader(http.StatusNotFound)を呼び出してユーザにHTTPエラーを報告する
- URLとハンドラの間の関連づけを簡単にするために、リクエストマルチプレクサ(ServeMux) net/httpは提供する
*/
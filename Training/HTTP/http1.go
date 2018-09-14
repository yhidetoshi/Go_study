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
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

/*
ブラウザで　localhost:8000

shoes: $50.00
socks: $5.00
*/
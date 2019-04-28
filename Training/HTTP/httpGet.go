package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://google.co.jp"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}

/*
- package http ( client.go )
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

- 参考
> https://qiita.com/taizo/items/c397dbfed7215969b0a5
*/
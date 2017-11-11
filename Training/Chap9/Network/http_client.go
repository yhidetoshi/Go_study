package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("http://golang.jp/hello.html")

	// Error Chekc
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status", response.Status)

	body, err := ioutil.ReadAll(response.Body)

	// Error Check
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output Body
	fmt.Println(string(body))
}

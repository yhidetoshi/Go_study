package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	url := "http://google.co.jp"
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

}

/*
package http ( request.go )
func NewRequest(method, url string, body io.Reader) (*Request, error)
*/
package main

import "fmt"

func main() {
	/*
	repositories = {
	  name = api-v1
	  name = api-v2
	}
	*/

	params := make(map[string]string)

	params["name"] = "api-v1"
	params["name"] = "api-v2"

	fmt.Println(params["name"]) // key名重複だから上書きされて作成されるリソースは api-v2だけ。
}

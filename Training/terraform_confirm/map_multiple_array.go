package main

import "fmt"

func main() {
	/*
	repositories = {
	    api1-v1 = {
	      name                 = "api-v1"
	      image_tag_mutability = "MUTABLE"
	      scan_on_push         = true
	    },
	    api-v2 = {
	      name                 = "api-v2"
	      image_tag_mutability = "IMMUTABLE"
	      scan_on_push         = false
	    },
	  }
	*/

	repositories := make(map[string]map[string]interface{})
	repositories["api-v1"] = make(map[string]interface{})
	repositories["api-v2"] = make(map[string]interface{})

	// api-v1
	repositories["api-v1"]["name"] = "api-v1"
	repositories["api-v1"]["image_tag_mutability"] = "MUTABLE"
	repositories["api-v1"]["scan_on_push"] = true

	// api-v2
	repositories["api-v2"]["name"] = "api-v2"
	repositories["api-v2"]["image_tag_mutability"] = "IMMUTABLE"
	repositories["api-v2"]["scan_on_push"] = false

	// output
	fmt.Println(repositories["api-v1"]["name"])
	fmt.Println(repositories["api-v1"]["image_tag_mutability"])
	fmt.Println(repositories["api-v1"]["scan_on_push"])
	fmt.Println()
	fmt.Println(repositories["api-v2"]["name"])
	fmt.Println(repositories["api-v2"]["image_tag_mutability"])
	fmt.Println(repositories["api-v2"]["scan_on_push"])
}

/* 実行結果
api-v1
MUTABLE
true

api-v2
IMMUTABLE
false
*/

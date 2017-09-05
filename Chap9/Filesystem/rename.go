package main

import (
	"fmt"
	"os"
)

func main() {
	// create Directory
	os.MkdirAll("testdir", 0777)

	// ディレクトリの名前変更
	err := os.Rename("testdir", "newdir")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

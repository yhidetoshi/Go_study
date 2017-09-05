package main

import (
	"os"
)

func main() {
	os.Mkdir("newdir", 0777)
}

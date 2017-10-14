package main

import(
	"encoding/binary"
	"crypto/rand"
	"strconv"
)

func genRand() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func main(){
	println(genRand())
}

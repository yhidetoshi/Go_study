package main

import (
	"io"
	"os"
	"log"
)

func main(){

	f_scr_name := os.Args[0]
	f_dst_name := os.Args[1]

	src, err := os.Open(f_scr_name)
	if err != nil{
		log.Println(err)
	}
	defer src.Close()

	dst, err := os.Create(f_dst_name)
	if err != nil{
		log.Println(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil{
		log.Println(err)
	}
}

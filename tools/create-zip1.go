// 圧縮対象のファイル名と圧縮名のファイル名を固定した場合

package main

import (
	"archive/zip"
//	"fmt"
	"io"
	"os"
)

func main() {

	var name string
	name = "output.txt"
	/*
	if len(os.Args) < 2 {
		fmt.Println("Usage: %s [zipname] [src] [src2] ...")
		return
	}
	*/
	dest, err := os.Create("output-txt.zip")
	//fmt.Println(os.Args[1])
	//dest, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	//for _, s := range os.Args[2:] {
	//for _, s := range name {
	if err := addToZip(name, zipWriter); err != nil {
		panic(err)
	}
	//}
	//fmt.Println(os.Args[2]) //output.txt
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	writer, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, src)
	if err != nil {
		return err
	}

	return nil
}

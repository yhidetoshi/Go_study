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
	dest, err := os.Create("output-txt.zip")

	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(dest)
	defer zipWriter.Close()

	if err := makeZip(name, zipWriter); err != nil {
		panic(err)
	}
}

func makeZip(filename string, zipWriter *zip.Writer) error {
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

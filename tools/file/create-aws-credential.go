package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"log"
	"archive/zip"
	"io"
)

const (
	N = 10 							// 読み込み行数の最大値
	INFILE = "./credentials.csv"    // 読み込みcsvファイル名を指定
)

// エラー処理
func fileError(err error){
	if err != nil {
		log.Fatal("Error", err)
	}
}


// CSVファイルを作成する
func outputFiles(_lineIn int, _buffIn [N]string){
	var lineOut int
	var outFilename string
	var outFilenameZip string

	// 2行目以降の行数だけファイルを作成(ユーザ毎にファイル作成)
	for lineOut = 0; lineOut < _lineIn; lineOut++ {

		if lineOut != 0 {
			outFilename = "./output"+strconv.Itoa(lineOut)+".csv"
			file, err := os.Create(outFilename)
			file.Write(([]byte)(_buffIn[0]+"\n"))
			file.Write(([]byte)(_buffIn[lineOut]))

			fileError(err)
			defer file.Close()

			// File(csv)をzip化する
			outFilenameZip = outFilename+".zip"
			outFilenameZip, err := os.Create(outFilenameZip)
			fileError(err)

			zipWriter := zip.NewWriter(outFilenameZip)
			defer zipWriter.Close()

			if err := makeZip(outFilename, zipWriter); err != nil {
				panic(err)
			}
		}

	}

}

func main() {
	var fp *os.File
	var err error
	var lineIn int
	var buffIn [N]string

	// Input File
	if len(os.Args) < 1 {
		fp  = os.Stdin
	} else {
		fp, err = os.Open(INFILE)
		fileError(err)
		defer fp.Close()
	}

	// 1行ずつファイルを読み込む
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		buffIn[lineIn] = scanner.Text()
		fmt.Println(scanner.Text())
		lineIn ++ //行数インクリメント
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	// ファイル行数と行毎のデータを配列で渡す
	outputFiles(lineIn, buffIn)
}

// File名を渡してzip化する
func makeZip(filename string, zipWriter *zip.Writer) error{

	src, err := os.Open(filename)
	fileError(err)
	defer src.Close()

	writer, err := zipWriter.Create(filename)
	fileError(err)

	_, err = io.Copy(writer, src)
	fileError(err)

	return nil
}

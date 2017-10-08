package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"log"
)
const (
	N = 10
	INFILE = "./credentials.csv"
)

func fileError(err error){
	if err != nil {
		log.Fatal("Error", err)
	}
}



func outputFiles(_lineIn int, _buffIn [N]string){
	var lineOut int
	var outFilename string

	for lineOut = 0; lineOut < _lineIn; lineOut++ {

		if lineOut != 0 {
			outFilename = "output"+strconv.Itoa(lineOut)+".csv"
			file, err := os.Create(outFilename)
			file.Write(([]byte)(_buffIn[0]+"\n"))
			file.Write(([]byte)(_buffIn[lineOut]))

			fileError(err)
			defer file.Close()
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

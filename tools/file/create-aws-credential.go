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

func main() {
	var fp *os.File
	var err error
	var lineIn, lineOut int
	var buffIn [N]string
	var outFilename string

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

	// 2行目以降をファイル出力する
	for lineOut = 0; lineOut < lineIn; lineOut++ {

		if lineOut != 0 {
			outFilename = "output"+strconv.Itoa(lineOut)+".csv"
			file, err := os.Create(outFilename)
			file.Write(([]byte)(buffIn[0]+"\n"))
			file.Write(([]byte)(buffIn[lineOut]))

			fileError(err)
			defer file.Close()
		}
	}
}

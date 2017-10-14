package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"log"
	//"archive/zip"
	//"io"
	"github.com/yeka/zip"
	//"bytes"
	//"io"
	//"bytes"
	"io"
	"encoding/binary"
	"crypto/rand"
)

const (
	N = 10 							// 読み込み行数の最大値
	INFILE = "./credentials.csv"    // 読み込みcsvファイル名を指定
	PWFILE = "zip-password.txt"
)

func fileError(err error){
	if err != nil {
		log.Fatal("Error", err)
	}
}

func genRand() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func createPWFile(pw string){
	file, err := os.Create(PWFILE)
	fileError(err)
	defer file.Close()

	file.Write(([]byte)(pw))
}

func outputFiles(_lineIn int, _buffIn [N]string){
	var lineOut int
	var outFilename string
	var outFilenameZip string

	// パスワード付きzipのパスワードファイルを生成
	password := genRand()
	createPWFile(password)

	for lineOut = 0; lineOut < _lineIn; lineOut++ {

		if lineOut != 0 {
			outFilename = "./output"+strconv.Itoa(lineOut)+".csv"
			file, err := os.Create(outFilename)
			file.Write(([]byte)(_buffIn[0]+"\n"))
			file.Write(([]byte)(_buffIn[lineOut]))

			fileError(err)
			defer file.Close()

			outFilenameZip = outFilename+".zip"
			outFilenameZip, err := os.Create(outFilenameZip)
			fileError(err)

			zipWriter := zip.NewWriter(outFilenameZip)
			defer zipWriter.Close()

			if err := makeZip(outFilename, zipWriter, password); err != nil {
				panic(err)
			}
		}
	}
}


func makeZip(filename string, zipWriter *zip.Writer, pw string) error {

	src, err := os.Open(filename)
	fileError(err)
	defer src.Close()

	dest, err := zipWriter.Encrypt(filename, pw, zip.AES256Encryption)
	fileError(err)

	_, err = io.Copy(dest, src)
	fileError(err)

	return nil
}


func main() {
	var fp *os.File
	var err error
	var lineIn int
	var buffIn [N]string

	if len(os.Args) < 1 {
		fp  = os.Stdin
	} else {
		//fp, err = os.Open(INFILE)
		fp, err = os.Open(os.Args[1])
		fileError(err)
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		buffIn[lineIn] = scanner.Text()
		fmt.Println(scanner.Text())
		lineIn ++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	outputFiles(lineIn, buffIn)
}



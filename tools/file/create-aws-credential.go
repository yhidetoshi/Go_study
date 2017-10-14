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
	//"bytes"
	//"io"
	//"bytes"
	"io"
)

const (
	N = 10 							// 読み込み行数の最大値
	INFILE = "./credentials.csv"    // 読み込みcsvファイル名を指定
	PW = "hoge"
)

func fileError(err error){
	if err != nil {
		log.Fatal("Error", err)
	}
}

func outputFiles(_lineIn int, _buffIn [N]string){   
	var lineOut int
	var outFilename string
	var outFilenameZip string

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

			if err := makeZip(outFilename, zipWriter); err != nil { 
				panic(err)
			}
		}
	}
}


func makeZip(filename string, zipWriter *zip.Writer) error {

	password := PW

	src, err := os.Open(filename)
	fileError(err)
	defer src.Close()

	dest, err := zipWriter.Encrypt(filename, password, zip.AES256Encryption)
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
		fp, err = os.Open(INFILE)
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



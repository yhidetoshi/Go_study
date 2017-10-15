package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	N       = 10                  // 読み込み行数の最大値
	INFILE  = "./credentials.csv" // 読み込みcsvファイル名を指定
	OUTFILE = "outCredentials_"
	CSV     = ".csv"
	ZIP     = ".zip"
	PWFILE  = "zip-password.txt"
)

/* Error処理 */
func fileError(err error) {
	if err != nil {
		log.Fatal("Error", err)
	}
}

/* zipのパスワードを発行 */
func genPWrand() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

/* zipパスワードファイルを作成 */
func createPWFile(pw string) {
	file, err := os.Create(PWFILE)
	fileError(err)
	defer file.Close()

	file.Write(([]byte)(pw))
}

func outputFiles(_lineIn int, _buffIn [N]string) {
	var lineOut int           // 出力ファイルの行数カウンタ
	var outFilename string    // ユーザ毎に分けた個別ファイル名(csv)
	var outFilenameZip string // ユーザ毎に分けた個別ファイル名(zip)

	password := genPWrand() // ランダムPWの生成
	createPWFile(password)  // PWを記入したテキストファイルを生成

	for lineOut = 0; lineOut < _lineIn; lineOut++ {

		if lineOut != 0 {
			outFilename = "./" + OUTFILE + strconv.Itoa(lineOut) + CSV
			file, err := os.Create(outFilename) // func Create(name string) (*File, error)
			file.Write(([]byte)(_buffIn[0] + "\n"))
			file.Write(([]byte)(_buffIn[lineOut]))

			fileError(err)
			defer file.Close()

			outFilenameZip = outFilename + ZIP
			outFilenameZip, err := os.Create(outFilenameZip) // func Create(name string) (*File, error)
			fileError(err)

			zipWriter := zip.NewWriter(outFilenameZip)
			defer zipWriter.Close()

			if err := makeZip(outFilename, zipWriter, password); err != nil {
				panic(err)
			}
		}
	}
}

/* PW付きZipを作成する */
func makeZip(filename string, zipWriter *zip.Writer, pw string) error {

	src, err := os.Open(filename)
	fileError(err)
	defer src.Close()

	dest, err := zipWriter.Encrypt(filename, pw, zip.AES256Encryption) // PW付きzipにする
	fileError(err)

	_, err = io.Copy(dest, src)
	fileError(err)

	return nil
}

func main() {
	var fp *os.File
	var err error
	var lineIn int       // 読み込むファイル(csv)の行数カウンタ
	var buffIn [N]string // 行毎のデータを格納用

	if len(os.Args) < 1 { // os.Args[0]はファイル読込み
		fp = os.Stdin
	} else {
		//fp, err = os.Open(INFILE)
		fp, err = os.Open(os.Args[1]) // func Open(name string) (*File, error)
		fileError(err)
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp) // func NewScanner(r io.Reader) *Scanner
	for scanner.Scan() {            // func (s *Scanner) Scan() bool
		buffIn[lineIn] = scanner.Text() // func (s *Scanner) Text() string
		fmt.Println(scanner.Text())
		lineIn++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	outputFiles(lineIn, buffIn)
}

![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Go_study/Golang-top.png)

## Go言語 やっていること



- 環境　
  - OS: 
    - MacOS
  - Goバージョン:  
    - go1.9
  - IDE: 
    - Gogland (https://www.jetbrains.com/go/)
- 学習
  - A Tour of Go (https://go-tour-jp.appspot.com/welcome/1)
  - 基礎からわかる Go言語(書籍)
  - Goならわかるシステムプログラミング(書籍)
  - プログラミング言語Go(書籍)
  - Paiza
  - Tools作り (https://github.com/yhidetoshi/Go_study/tree/master/tools)
  - aws-sdk-goを使ったツール作り
  - aws-sdk-goとFlagを使ったCLIツール作成(https://github.com/yhidetoshi/clitoolgoaws)
  - PointerのMemo (https://github.com/yhidetoshi/Go_study/tree/master/Training/Pointer/README.md)
  - StructのMemo (https://github.com/yhidetoshi/Go_study/blob/master/Training/Struct/README.md)
- Goパス(ex):
  - ```export GOPATH=`pwd` ```
  - `$go env`
    
## Goのインストール
- macOS
  - `brew install go`
  - `go version`

- MacOS上でコーディングしたコードを Amazon Linux上で実行するためにクロスコンパイルする
  - `$ GOOS=linux GOARCH=amd64 go build {FILENAME}`

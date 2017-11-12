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
    - 基礎からわかる Go言語
    - Paiza
    - Tools作り
    - aws-sdk-goを使ったツール作り
    - aws-sdk-goとFlagを使ったCLIツール作成(https://github.com/yhidetoshi/clitoolgoaws)
    - PointerのMemo (https://github.com/yhidetoshi/Go_study/tree/master/Training/Pointer/README.md)
    - StructのMemo (https://github.com/yhidetoshi/Go_study/blob/master/Training/Struct/README.md)
  - Gogland(pkg)のパス(ex):
    - `$HOME/go/src/github.com/`
  - Go(pkg)のパス(ex):
    - `$HOME/.go/src`
    
## Goのインストール
- macOS
  - `brew install go`
  - `go version`

- MacOS上でコーディングしたコードを Amazon Linux上で実行するためにクロスコンパイルする
  - `$ GOOS=linux GOARCH=amd64 go build {FILENAME}`

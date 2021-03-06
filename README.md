![Alt Text](https://github.com/yhidetoshi/Pictures/raw/master/Go_study/Golang-top.png)

## Go言語 やっていること



- 環境　
  - OS:
    - MacOS
  - Goバージョン:  
    - go1.9
  - IDE:
    - Gogland (https://www.jetbrains.com/go/) ※ Goglandが無料で使えた時に使用
    - Goland
    - Atom + Go plugin
- 学習&開発
  - http://golang.jp/go_spec
  - https://golang.org/
    - https://golang.org/pkg/
  - Effective Go
    - https://golang.org/doc/effective_go.html
  - Goのコード(GitHub)
    - https://github.com/golang/go/tree/master/src/strings
  - A Tour of Go (https://go-tour-jp.appspot.com/welcome/1)
  - 基礎からわかる Go言語(書籍)
  - Goならわかるシステムプログラミング(書籍)
  - プログラミング言語Go(書籍)
  - Go言語による並行処理(書籍)
  - Paiza
  - Tools作り (https://github.com/yhidetoshi/Go_study/tree/master/tools)
  - aws-sdk-goを使ったツール作り
    - (ex) AMIの自動削除(最新のN個を残す & ssmパラメータAMIは対象外)
      - https://github.com/yhidetoshi/Go_study/blob/master/aws-go/ec2Ami.go
      - https://github.com/yhidetoshi/GoAWSDeleteAmisLaunchConfigsTool
  - `ore-aws`(aws-sdk-goとFlagを使ったCLIツール作成)
    - https://github.com/yhidetoshi/go-awscli-tool
  - Mackerel (Plugin/Tool) 開発
    - https://github.com/yhidetoshi/mackerel-plugin-awsbilling 
    - https://github.com/yhidetoshi/mackerelCPUAlertTool
    - https://github.com/yhidetoshi/mackerelCheckAutoRetireFail
  - Cloud Function (Google Cloud)
    - https://github.com/yhidetoshi/gcloud-function-examples
  - Serverless-Lambda(GO)の学習
    - https://github.com/yhidetoshi/serverless-dev-go
    - https://github.com/yhidetoshi/Golang-Lambda
  - GolangCI
    - https://github.com/yhidetoshi/GolangCI-Dev
  - NatureRemo-Mackerel可視化連携
    - https://github.com/yhidetoshi/natureRemo-mackerel
  - Slack
    - https://github.com/yhidetoshi/Go_study/tree/master/Training/Slack
  - Goパス(ex):
    - ```export GOPATH=`pwd` ```
    - `$go env`

## Goのインストール
- macOS
  - `brew install go`
  - `go version`
  - GOPATHを通す

- MacOS上でコーディングしたコードを Amazon Linux上で実行するためにクロスコンパイルする
  - `$ GOOS=linux GOARCH=amd64 go build -o {FREE_FILENAME} {TARGET}.go`
- Windows用
  - `GOOS=windows GOARCH=amd64 go build -o {FREE_FILENAME} {TARGET}.go`

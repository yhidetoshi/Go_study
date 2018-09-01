# Interfaceについてのめも


## Interfaceの種類
- 参考
  - http://jxck.hatenablog.com/entry/20130325/1364251563
- type
  - interface (interfaceDocument.go)
  - mixin (interfaceDocumentMixin.go)
  - DuckType (interfaceDocumetDuckType.go)
  - overide (interfaceDocumentOverride.go)
  - interface型
  - Generics (interfaceDocumentGenerics.go)
  - interfaceの設計
  - tag


## fmtパッケージ

### String()
- type Stringer

- https://github.com/golang/go/blob/master/src/fmt/print.go
  - go言語のgithubで定義している箇所
```
type Stringer interface{
  String() string
}
```
- func(b *Builder) String() string


- https://golang.org/pkg/fmt/#Print
  - func Print
    - `func Print(a ...interface{}) (n int, err error)`
      - Printは interface型を受けるので、どんなものを受け取る
  - func Printf
    - `func Printf(format string, a ...interface{}) (n int, err error)`
      - formatで受けるので、 %dや%vなど指定してから値を受ける

- fmtパッケージのPrint関数
```
// 標準出力へ書き出す関数
func Print(a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)

// ファイル（io.Writer）へ書き出す関数
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

// 結果を string 値として返す関数
func Sprint(a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string

// エラーを生成する関数
func Errorf(format string, a ...interface{}) error
```
- `→ fmt.Prinln , fmt.Printf("%d", val)`


- https://golang.org/src/strings/builder.go
```
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}
```

- https://golang.org/pkg/strings/#Builder
  - A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.
```
type Builder struct {
        // contains filtered or unexported fields
}
```


## osパッケージ
- func Open
  - func Open(name string) (*File, error)
    - Open opens the named file for reading. If successful, methods on the returned file can be used for reading
      - file, err := os.Open(`/path/to/file`)

- (ex)    
```
file, err := os.Open(`/path/to/file`)
if err != nil {
  // open Error処理
}
defer file.Close()
```

## io/ioutilパッケージ

- func (*File) Read
  - func (f *File) Read(b []byte) (n int, err error)
    - Read reads up to len(b) bytes from the File. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.


```
type Reader interface {
        Read(p []byte) (n int, err error)
}

* 引数である p は、読み込んだ内容を一時的に入れておくバッファ
```

- (ex)
```
content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
```

## プログラミング言語Goより

### インターフェース
- 特徴
  - インターフェースは暗黙的に満足される
    - 具象型が満足するすべてのインタフェースを宣言する必要がない
    - 単に必要なメソッドを持っているだけで十分

- 契約としてのインタフェース

- **具象型**
  - その値の正確な表現を決めていて、数値に対する演算、スライスに対するインデックス指定やappend, rangeなどのような、その表現の本質的な操作を公開している

- **抽象型**
  - インターフェース型
    - インターフェースはその値の表現つまり内部構造を公開していないし、その値がサポートする基本操作の集まりも公開していない。
    - インターフェースは値が持つメソッドのいくつかを示しているだけ。
    - その値が何であるかについては何もわからない。その値で何ができるかを知っているだけ

- 標準出力へ結果を書き出す: `fmt.Printf`
- stringとして結果を返す: `fmt.Sprintf
  - → 結果をフォーマットするために、難しい部分を複製しなければならないのは困難
  - インターフェースのおかげで、複製せずに簡単に利用できる

```
package fmt
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
```

```
package io
type Writer interface {
  Write(p []byte) (n int, err error)
}
```
→ io.Writer インターフェースはFprintfとその呼び出し元との間の契約を定義している

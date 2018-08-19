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

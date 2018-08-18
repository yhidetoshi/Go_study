# Interfaceについてのめも


## fmtパッケージ

### String()
- type Stringer

```
type Stringer interface{
  String() string
}
```
- func(b *Builder) String() string


- https://golang.org/src/strings/builder.go
```
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}
```

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

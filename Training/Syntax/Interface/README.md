# Interfaceについてのめも

## osパッケージ
- func Open
  - func Open(name string) (*File, error)
    - Open opens the named file for reading. If successful, methods on the returned file can be used for reading
      - file, err := os.Open(`/path/to/file`)  

```
file, err := os.Open(`/path/to/file`)
if err != nil {
  // open Error処理
}
defer file.Close()
```

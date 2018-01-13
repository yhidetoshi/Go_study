# 　使った構文Memo

- **標準入力**
```
var num int
fmt.Scan(&num)
```
- `strings.Repeat(somethingString, Number)`
  - `Number` の分だけ `somethingString`　を繰り返す
```
fmt.Println(strings.Repeat(aster, num))
```

- `strings.Contains(ch_str, in_str)`
  - `ch_str`の文字列に `in_str` が含まれているかを判定（真偽値）

# 　使った構文Memo

- **標準入力その１**
```
var num int
fmt.Scan(&num)
```

- **標準入力その2**
```
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}


input = StrStdin()
```

- `strings.Repeat(somethingString, Number)`
  - `Number` の分だけ `somethingString`　を繰り返す
```
fmt.Println(strings.Repeat(aster, num))
```

- `strings.Contains(ch_str, in_str)`
  - `ch_str`の文字列に `in_str` が含まれているかを判定（真偽値）

- `strings.Replace(strBefore, "False", "True", -1)`
- `strings.Replace(strBefore, "at", "@", -1)`

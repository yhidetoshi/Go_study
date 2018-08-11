// https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8 から学習
// can't run

type Person intarface {
  Title () string
  Name() string
}

type person struct{
  firstName string
  lastName string
}

func (p *person) Name() string{
  return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

/*
以下の*person型は，Nameメソッドを持っています．
しかし，Titleメソッドは持っていないため，
Personインタフェースを実装していることにはなりません．


しかし，*person型をTitleメソッドを実装した構造体に埋め込むことで，
Personインタフェースを実装することができます．
以下の例では，*female型と*male型にTitleメソッドを実装させています．
それぞれの型には*person型が埋め込まれているため，
Nameメソッドも実装していることになります．
そのため，*female型と*male型はPersonインタフェースを実装していることになります．
*/

type Gender int

const (
  Female = iota
  Male
)

type female struct {
  *person
}

func (f *female) Title() string{
  return "Ms."
}

type male struct {
  *person
}

func (m *male) Title() string{
  return "Mr."
}

/*
Go言語では，小文字で始まる型はそのパッケージ内からしか使用することができません．
そのため，female型とmale型は他のパッケージから隠蔽されています．
他のパッケージからは，これらの型を意識せずPersonインタフェースとして
使用できることが好ましいでしょう．

Personインタフェースを実装した値を返す，
NewPerson関数を考えてみましょう．引数に性別(gender)を指定することで，
内部で生成する構造体の型を切り替えています．このように呼び出し元では，
Personインタフェースを実装しているのが*female型なのか，
*male型のなのか意識せず，Titleメソッドの実装を切り替えることができています．
*/



/*
Personインタフェースを実装した値を返す，
NewPerson関数を考えてみましょう．引数に性別(gender)を指定することで，
内部で生成する構造体の型を切り替えています．このように呼び出し元では，
Personインタフェースを実装しているのが*female型なのか，
*male型のなのか意識せず，Titleメソッドの実装を切り替えることができています．
*/

func NewPerson(gender Gender, firstName, lastName string) Person {
  p := &person{firstName, lastName}

  if gender == Female {
    return &female{p}
  }
  return &male{p}
}

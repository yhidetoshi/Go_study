package main

import "fmt"

func main(){
  var pointer *int // アドレスを格納できるようにしている [*]
  var n int = 100  
  pointer = &n


  fmt.Println("nのアドレス", &n)       // nのアドレス 0xc420016090 <アドレスの場所を指している>
  fmt.Println("pointerの値", pointer) // pointerの値 0xc420016090 <アドレスの場所を表示している>

  fmt.Println("nの値",n)                 // nの値 100
  fmt.Println("pointerの中身", *pointer) //pointerの中身 100

}

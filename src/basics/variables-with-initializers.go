package main

import "fmt"

// var 宣言では、変数毎に初期化子( initializer )を与えることができます。
var i, j int = 1, 2

func main() {
  // 初期化子が与えられている場合、型を省略できます。その変数は初期化子が持つ型になります。
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}

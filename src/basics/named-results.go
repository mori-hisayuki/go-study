package main

import "fmt"

// 戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。
func split(sum int) (x, y int) {
  x = sum * 4 /9
  y = sum - x
  // 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。
  return
}

func main() {
  fmt.Println(split(17))
}

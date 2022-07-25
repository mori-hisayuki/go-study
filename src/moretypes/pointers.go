package main

import "fmt"

func main() {
  i, j := 42, 2701

  // 変数iのメモリアドレス取得
  p := &i
  // 変数iのメモリアドレスから変数の値を取得
  fmt.Println(*p)
  // メモリアドレスを通して変数iに値を代入
  *p = 21
  // 変数iの中身を表示
  fmt.Println(i)

  // 変数jのメモリアドレスを取得
  p = &j
  // メモリアドレスを通して変数jに計算結果を代入
  // j = j/37と同位
  *p = *p / 37
  // 変数jの中身を表示
  fmt.Println(j)
}

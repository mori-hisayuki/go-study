package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  // 変数vを表示
  fmt.Println(v)
  // ポインタを通して変数vを表示
	fmt.Println(*p)
	fmt.Println(p)

}

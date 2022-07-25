package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  // structのフィールドは、ドット( . )を用いてアクセスします。
  v.X = 4
  fmt.Println(v.X)
}

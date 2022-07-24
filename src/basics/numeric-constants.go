package main

import "fmt"

// 数値の定数は、高精度な 値 ( values )です
// 型のない定数は、その状況によって必要な型を取ることになります。
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)

func needInt(x int) int {
  return x*10 + 1
}

func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}

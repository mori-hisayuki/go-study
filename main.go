package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
  var help = false
  var format = ""
  flag.BoolVar(&help, "h", false, "show help")
  flag.StringVar(&format,"f","%s\n", "set format")
  flag.Parse()

  if help {
    flag.Usage()
    return
  }

  for i, arg := range flag.Args() {
    fmt.Print(strconv.FormatInt(int64(i), 10) + ":")
    fmt.Print(strconv.Itoa(i) + ":")
    fmt.Print("%d", i)
    fmt.Print(format, arg)
  }
}

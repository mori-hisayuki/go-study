package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var help = false
	var format = ""
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&format,"f", "%s\n", "set format")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

    dict := map[string][]string {
		"kanji": {"零", "壱", "弐", "参", "肆"},
		"roma": {"?", "Ⅰ", "Ⅱ", "Ⅲ", "Ⅳ"},
	}


	for _, arg := range flag.Args() {
		n, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Fprintln(os.Stderr, "数値のパースに失敗しました:" +err.Error())
			os.Exit(1)
		}
		if n < 0 || n > 4 {
			fmt.Fprintln(os.Stderr, "数値が範囲外です。0~4にしてください: " +arg)
			os.Exit(1)
		}
		if dict[format] == nil {
			fmt.Fprintln(os.Stderr, "kanji か roma を指定してね" +format)
			os.Exit(1)
		}
		fmt.Println(dict[format][n])
	}
}

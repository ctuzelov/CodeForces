package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var res strings.Builder
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve() {
	var n int
	fmt.Fscan(in, &n)
	if n&1 == 1 {
		for i := 0; i < n; i++ {
			res.WriteString("69 ")
		}
	} else {
		res.WriteString("1 3 ")
		for i := 2; i < n; i++ {
			res.WriteString("2 ")
		}
	}
	res.WriteString("\n")
}

func main() {
	var t int
	fmt.Fscan(in, &t)
	defer out.Flush()

	for i := 0; i < t; i++ {
		solve()
	}

	fmt.Fprintln(out, res.String())
}

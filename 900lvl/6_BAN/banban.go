package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t   int
		res strings.Builder
	)

	fmt.Fscan(in, &t)
	for t > 0 {
		var n int
		fmt.Fscan(in, &n)

		if n%2 == 0 {
			res.WriteString(strconv.Itoa(n / 2))
		} else {
			res.WriteString(strconv.Itoa(n/2+1) + "\n")
			res.WriteString(strconv.Itoa(3) + " " + strconv.Itoa((n*3+1)/2))
		}

		res.WriteString("\n")

		q, k := 1, n*3
		for q < k-3 {
			res.WriteString(strconv.Itoa(q) + " " + strconv.Itoa(k) + "\n")
			q += 3
			k -= 3
		}
		t--
	}

	fmt.Fprintln(out, res.String())
}

/*

*/

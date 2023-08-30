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
	var t int
	var s int
	var res strings.Builder
	fmt.Fscan(in, &s)

	for j := 0; j < s; j++ {
		fmt.Fscan(in, &t)
		if (t == 3 || t == 1) && t < 5 {
			res.WriteString("NO" + "\n")
			continue
		}

		res.WriteString("YES" + "\n")

		if t%2 == 0 {
			for j := 0; j < t; j++ {
				if (j+1)%2 != 0 {
					res.WriteString("1" + " ")
				} else {
					res.WriteString("-1" + " ")
				}
			}
			res.WriteString("\n")
			continue
		}

		n := (-t + 3) / 2
		for i := 0; i < t; i++ {
			if (i+1)%2 != 0 {
				res.WriteString(strconv.Itoa(n) + " ")
			} else {
				res.WriteString(strconv.Itoa(-n+1) + " ")
			}
		}
		res.WriteString("\n")
	}
	fmt.Fprintln(out, res.String())
}

/*

	1 -1 1 -1 1 -1 1 -1

	-6 7 -6 7 -6 7 -6 7 -6 7 -6 7 -6 7 -6

	-5 6 -5 6 -5 6 -5 6 -5 6 -5 6 -5

	-4 5 -4 5 -4 5 -4 5 -4 5 -4

	-3 4 -3 4 -3 4 -3 4 -3

	-2 3 -2 3 -2 3 -2

	-1 2 -1 2 -1

	5 = -1; 7 := -2; 9 = -3

	-n/2 + 3/2
*/

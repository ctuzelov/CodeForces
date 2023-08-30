package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N = 1e5 + 5

func main() {
	var t int
	var res strings.Builder
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &t)
	for t > 0 {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		sort.Ints(a)

		if a[0] == a[n-1] {
			res.WriteString(strconv.FormatInt(int64(n)*int64(n-1), 10) + "\n")
			t--
			continue
		}

		mn := 0
		mx := n - 1

		for a[0] == a[mn] {
			mn++
		}

		for a[n-1] == a[mx] {
			mx--
		}

		l := int64(mn)
		r := int64(n - mx - 1)

		res.WriteString(strconv.FormatInt(2*l*r, 10) + "\n")

		t--
	}

	fmt.Fprintln(out, res.String())
}

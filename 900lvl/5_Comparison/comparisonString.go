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

	for i := 0; i < t; i++ {
		var s int
		var input string
		count, ans := 1, 1
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &input)

		for j := 0; j < s-1; j++ {
			if input[j] == input[j+1] {
				count++
			} else {
				count = 1
			}
			ans = Max(ans, count)
		}
		res.WriteString(strconv.Itoa(ans+1) + "\n")
	}

	fmt.Fprintln(out, res.String())
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*

4
4
<<>>
0

4
>><<
5
>>>>>
7
<><><><



*/

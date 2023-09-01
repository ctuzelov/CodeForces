package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t   int
		res string
	)
	fmt.Fscan(in, &t)

	for t > 0 {

		var slv string
		fmt.Fscan(in, &slv)

		sign := slv[1]
		n1, _ := strconv.Atoi(string(slv[0]))
		n2, _ := strconv.Atoi(string(slv[2]))

		if sign == '-' {
			res += strconv.Itoa(n1 - n2)
		} else {
			res += strconv.Itoa(n1 + n2)
		}

		res += "\n"

		t--
	}

	fmt.Fprintln(out, res)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rows int
	var res string
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &rows)

	for i := 0; i < rows; i++ {
		size, k, cml := 0, 0, 1
		fmt.Fscan(in, &size)
		fmt.Fscan(in, &k)
		m := make([]int, size)

		for j := 0; j < size; j++ {
			fmt.Fscan(in, &m[j])
		}

		for j := 0; j < size; j++ {
			if m[j] == cml {
				cml++
			}
		}

		res += strconv.Itoa((size-cml+k)/k) + "\n"
	}
	fmt.Fprintln(out, res)
}

/*

	2,3,1,4

	k = 2
	3,5,1,8,4,2

	5,1,8,2,3,4
	1,2,3,4,5,8

*/

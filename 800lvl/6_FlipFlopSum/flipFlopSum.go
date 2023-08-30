package main

import (
	"fmt"
	"strconv"
)

func main() {

	var res string
	var rows int
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		l := 0
		fmt.Scan(&l)
		neg, acc := false, true
		sum := 0
		m := make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&m[j])
			sum += m[j]
		}

		for j := 0; j < l; j++ {
			if m[j] == -1 {
				neg = true
				if j+1 < l && m[j+1] == -1 {
					res += strconv.Itoa(sum+4) + "\n"
					acc = false
					break
				}
			}
		}

		if acc && neg {
			res += strconv.Itoa(sum) + "\n"
		} else if acc && !neg {
			res += strconv.Itoa(sum-4) + "\n"
		}

	}

	fmt.Println(res[:len(res)-1])

}

/*
	1 1 1 1 1
	-1 -1 1 1 1

	-1 1 1 1 1
	1 -1 1 1 1

	-1 1 1 -1 -1
	-1 1 1 1 1

*/

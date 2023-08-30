package main

import (
	"fmt"
	"strconv"
)

func main() {
	var res string
	var t int
	fmt.Scan(&t) // Number of test cases

	for i := 0; i < t; i++ {
		m, n := 0, 0
		fmt.Scan(&n) // Number of elements in the list

		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j]) // List of elements
		}

		for j := 0; j < n; j++ {
			if a[j] == 2 {
				m++
			}
		}
		if m%2 != 0 {
			res += strconv.Itoa(-1) + "\n"
			continue
		} else if m == 0 {
			res += strconv.Itoa(1) + "\n"
			continue
		}

		count := 0

		for j := 0; j < n; j++ {
			if a[j] == 2 {
				count++
				m--
			}
			if count == m {
				res += strconv.Itoa(j+1) + "\n"
				break
			}
		}
	}
	fmt.Println(res)
}

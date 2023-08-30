package main

import (
	"fmt"
	"strconv"
)

func main() {
	var rows int
	var num string
	var res string
	pi := "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		fmt.Scan(&num)
		j := 0
		for j < len(num) && num[j] == pi[j] {
			j++
		}
		res += strconv.Itoa(j) + "\n"
	}
	fmt.Println(res[:len(res)-1])
}

package main

import (
	"fmt"
	"strconv"
)

var (
	rows int
	res  string
)

func main() {
	fmt.Scan(&rows)
	for i := 0; i < rows; i++ {
		item := 0
		fmt.Scan(&item)
		if item == 0 {
			res += "-1"
			continue
		}
		res += strconv.Itoa(item - 1)
		res += "\n"
	}
	fmt.Println(res)
}

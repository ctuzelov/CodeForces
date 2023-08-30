package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var res string
	var t int
	fmt.Scan(&t) // Number of test cases

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n) // Number of elements in the list

		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j]) // List of elements
		}

		sort.Ints(a)

		if a[0] == a[1] {
			res += "NO\n"
		} else {
			res += "YES\n"
			temp := append([]int{}, a[n-1], a[n-2])
			temp = append(temp, a[:n-2]...)
			res += conv(temp)
		}
	}
	fmt.Println(res)
}

func conv(arr []int) string {
	valuesText := []string{}
	for i := range arr {
		valuesText = append(valuesText, strconv.Itoa(arr[i]))
	}
	res := strings.Join(valuesText, " ")
	return res + "\n"
}

package main

import (
	"fmt"
	"strconv"
)

var (
	input  [][]int
	result string
)

func main() {
	GamingForces()
}

func Read() {
	var (
		rows int
		size int
		item int
	)
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		var temp []int
		fmt.Scan(&size)
		for j := 0; j < size; j++ {
			fmt.Scan(&item)
			temp = append(temp, item)
		}
		input = append(input, temp)
	}
	fmt.Println(input)
}

func GamingForces() {
	Read()
	for i := 0; i < len(input); i++ {
		count := 0
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 1 {
				count++
			}
		}
		result += strconv.Itoa(len(input[i]) - (count)/2)
		result += "\n"
	}
	fmt.Println(result)
}

/*

	1 1 1 2 2

*/

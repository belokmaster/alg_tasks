package main

import (
	"math"
	"sort"
)

func minOperationsTask(grid [][]int, x int) int {
	var matrix []int
	for _, row := range grid {
		matrix = append(matrix, row...)
	}
	sort.Ints(matrix)

	for i := 1; i < len(matrix); i++ {
		if (matrix[i]-matrix[0])%x != 0 {
			return -1
		}
	}

	target := matrix[len(matrix)/2]
	ans := 0
	for _, num := range matrix {
		ans += int(math.Abs(float64(num-target))) / x
	}

	return ans
}

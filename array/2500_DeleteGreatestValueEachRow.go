package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	ans := 0
	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		sort.Ints(grid[i])
	}

	// по столбцам
	for j := 0; j < m; j++ {
		maxElement := int(-1e9)

		// по строчкам
		for i := 0; i < n; i++ {
			if grid[i][m-1-j] > maxElement {
				maxElement = grid[i][m-1-j]
			}
		}

		ans += maxElement
	}

	return ans
}

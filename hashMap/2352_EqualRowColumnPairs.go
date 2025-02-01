package main

import "strconv"

// прикол с разделителями строк, чтобы не было проблем при строках 12 3 и 3 12.
func equalPairs(grid [][]int) int {
	n := len(grid)
	hashMap := make(map[string]int)
	ans := 0

	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			row += strconv.Itoa(grid[i][j]) + ","
		}
		hashMap[row]++
	}

	for j := 0; j < n; j++ {
		column := ""
		for i := 0; i < n; i++ {
			column += strconv.Itoa(grid[i][j]) + ","
		}
		if count, ok := hashMap[column]; ok {
			ans += count
		}
	}

	return ans
}
